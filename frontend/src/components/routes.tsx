/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import React from 'react';
import { Switch } from 'react-router-dom';
import { Section } from './misc/common';
import { Route, Redirect } from 'react-router';
import { PageComponentType, PageProps } from './pages/Page';
import TopicList from './pages/topics/Topic.List';
import TopicDetails from './pages/topics/Topic.Details';
import GroupList from './pages/consumers/Group.List';
import GroupDetails from './pages/consumers/Group.Details';
import { uiState } from '../state/uiState';
import { api } from '../state/backendApi';
import { CollectionIcon, FilterIcon } from '@heroicons/react/outline';
import { Feature, FeatureEntry, isSupported } from '../state/supportedFeatures';
import { UserPermissions } from '../state/restInterfaces';
import { AppFeature, AppFeatures } from '../utils/env';
import { AnimatePresence } from '../utils/animationProps';
import { NavLinkProps } from '@redpanda-data/ui/dist/components/Nav/NavLink';
import { TopicProducePage } from './pages/topics/Topic.Produce';

//
//	Route Types
//
type IRouteEntry = PageDefinition<any>;

export interface PageDefinition<TRouteParams = {}> {
    title: string;
    path: string;
    pageType: PageComponentType<TRouteParams>;
    routeJsx: JSX.Element;
    icon?: (props: React.ComponentProps<'svg'>) => JSX.Element;
    menuItemKey?: string; // set by 'CreateRouteMenuItems'
    visibilityCheck?: () => MenuItemState;
}


// Generate content for <Menu> from all routes
export function createVisibleSidebarItems(entries: IRouteEntry[]): NavLinkProps[] {
    return entries.map((entry) => {
        // Menu entry for Page
        if (entry.path.includes(':'))
            return null; // only root-routes (no param) can be in menu
        if (!entry.icon)
            return null; // items without icon do not appear in the sidebar

        let isEnabled = true;
        let disabledText: JSX.Element = <></>;
        if (entry.visibilityCheck) {
            const visibility = entry.visibilityCheck();
            if (!visibility.visible) return null;

            isEnabled = visibility.disabledReasons.length == 0;
            if (!isEnabled)
                disabledText = disabledReasonText[visibility.disabledReasons[0]];
        }
        const isDisabled = !isEnabled;

        return {
            title: entry.title as string,
            to: entry.path as string,
            icon: entry.icon as any,
            isDisabled: isDisabled as boolean,
            disabledText: disabledText as unknown as string,
        };
    }).filter(x => x != null && x != undefined) as NavLinkProps[];
}

// Convert routes to <Route/> JSX declarations
function EmitRouteViews(entries: IRouteEntry[]): JSX.Element[] {
    return entries.map(e => e.routeJsx);
}

export const RouteView = (() =>
    <AnimatePresence mode="wait">
        <Switch>
            {/* Index */}
            <Route exact path="/" render={() => <Redirect to="/topics" />} />

            {/* Emit all <Route/> elements */}
            {EmitRouteViews(APP_ROUTES)}

            <Route render={rp => {
                uiState.pageTitle = '404';
                return (
                    <Section title="404">
                        <div><h4>Path:</h4> <span>{rp.location.pathname}</span></div>
                        <div><h4>Query:</h4> <pre>{JSON.stringify(rp.location.search, null, 4)}</pre></div>
                    </Section>
                )
            }} />

        </Switch>
    </AnimatePresence>
)

enum DisabledReasons {
    'notSupported', // kafka cluster version too low
    'noPermission', // user doesn't have permissions to use the feature
    'enterpriseFeature'
}

const disabledReasonText: { [key in DisabledReasons]: JSX.Element } = {
    [DisabledReasons.noPermission]:
        <span>You don't have premissions<br />to view this page.</span>,
    [DisabledReasons.notSupported]:
        <span>The Kafka cluster does not<br />support this feature.</span>,
    [DisabledReasons.enterpriseFeature]:
        <span>This feature requires an enterprise license.</span>,
} as const;

interface MenuItemState {
    visible: boolean;
    disabledReasons: DisabledReasons[];
}

function MakeRoute<TRouteParams>(
    path: string,
    page: PageComponentType<TRouteParams>,
    title: string,
    icon?: (props: React.ComponentProps<'svg'>) => JSX.Element,
    exact: boolean = true, showCallback?: () => MenuItemState): PageDefinition<TRouteParams> {

    const route: PageDefinition<TRouteParams> = {
        title,
        path,
        pageType: page,
        routeJsx: (null as unknown as JSX.Element), // will be set below
        icon,
        visibilityCheck: showCallback,
    }

    // todo: verify that path and route params match
    route.routeJsx = <Route path={route.path} key={route.title} exact={exact ? true : undefined} render={rp => {
        const matchedPath = rp.match.url;
        const { ...params } = rp.match.params;

        if (uiState.currentRoute && uiState.currentRoute.path != route.path) {
            //console.log('switching route: ' + routeStr(ui.currentRoute) + " -> " + routeStr(route));
        }

        const pageProps: PageProps<TRouteParams> = {
            matchedPath,
            ...params,
        } as PageProps<TRouteParams>;

        uiState.currentRoute = route;
        return <route.pageType {...pageProps} />
    }} />;

    return route;
}

function routeVisibility(
    visible: boolean | (() => boolean),
    requiredFeatures?: FeatureEntry[],
    requiredPermissions?: UserPermissions[],
    requiredAppFeatures?: AppFeature[],
): () => MenuItemState {
    return () => {
        const v = typeof visible === 'boolean'
            ? visible
            : visible();

        const disabledReasons: DisabledReasons[] = [];
        if (requiredFeatures)
            for (const f of requiredFeatures) {
                if (!isSupported(f)) {
                    disabledReasons.push(DisabledReasons.notSupported);
                    break;
                }
            }

        if (requiredPermissions && api.userData)
            for (const p of requiredPermissions) {
                const hasPermission = api.userData[p];
                if (!hasPermission) {
                    disabledReasons.push(DisabledReasons.noPermission);
                    break;
                }
            }

        if (requiredAppFeatures) {
            for (const f of requiredAppFeatures)
                if (AppFeatures[f] == false) {
                    disabledReasons.push(DisabledReasons.enterpriseFeature);
                    break;
                }
        }

        return {
            visible: v,
            disabledReasons: disabledReasons
        }
    }
}

//
// Route Definitions
// If a route has one or more parameters it will not be shown in the main menu (obviously, since the parameter would have to be known!)
//
export const APP_ROUTES: IRouteEntry[] = [

    MakeRoute<{}>('/topics', TopicList, 'Topics', CollectionIcon),
    MakeRoute<{ topicName: string }>('/topics/:topicName', TopicDetails, 'Topics'),
    MakeRoute<{ topicName: string }>('/topics/:topicName/produce-record', TopicProducePage, 'Produce Record'),

    MakeRoute<{}>('/groups', GroupList, 'Consumer Groups', FilterIcon, undefined,
        routeVisibility(true, [Feature.ConsumerGroups])
    ),
    MakeRoute<{ groupId: string }>('/groups/:groupId/', GroupDetails, 'Consumer Groups'),

].filterNull();



