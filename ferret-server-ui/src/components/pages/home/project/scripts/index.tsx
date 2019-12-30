import React, { Fragment } from 'react';
import { Route, Switch } from 'react-router-dom';
import { Page, PageProps } from '../../../../common/page';

const LoadableListPage = React.lazy(() => import('./scripts') as any);
const LoadableDetailsPage = React.lazy(() => import('./script/index') as any);

export type Params = never;
export interface Props extends PageProps<Params> {
    project: string;
}

export default class ProjectScriptsPage extends Page<Params, Props> {
    constructor(props: Props) {
        super(props);

        this.__renderList = this.__renderList.bind(this);
        this.__renderDetails = this.__renderDetails.bind(this);
    }

    public render(): any {
        const { match } = this.props;

        return (
            <Fragment>
                <Switch>
                    <Route
                        path={match.path}
                        component={this.__renderList}
                        exact
                    />
                    <Route
                        path={`${match.path}/new`}
                        component={this.__renderDetails}
                    />
                    <Route
                        path={`${match.path}/:id`}
                        component={this.__renderDetails}
                    />
                </Switch>
            </Fragment>
        );
    }

    private __renderList(props: any): any {
        return <LoadableListPage {...props} projectId={this.props.project} />;
    }

    private __renderDetails(props: any): any {
        return (
            <LoadableDetailsPage {...props} projectId={this.props.project} />
        );
    }
}
