import Route from '@ember/routing/route';
import { inject as service } from '@ember/service';
import ApiService from 'waypoint/services/api';
import { Ref, GetBuildRequest } from 'waypoint-pb';
import { AppRouteModel } from '../app';

interface BuildModelParams {
  build_id: string;
}
export default class BuildDetail extends Route {
  @service api!: ApiService;

  breadcrumbs(model: AppRouteModel) {
    if (!model) return [];
    return [
      {
        label: model.application.application,
        icon: 'git-repository',
        args: ['workspace.projects.project.app'],
      },
      {
        label: 'Builds',
        icon: 'build',
        args: ['workspace.projects.project.app.builds'],
      },
    ];
  }

  async model(params: BuildModelParams) {
    // Setup the build request
    let ref = new Ref.Operation();
    ref.setId(params.build_id);
    let req = new GetBuildRequest();
    req.setRef(ref);

    let build = await this.api.client.getBuild(req, this.api.WithMeta());
    return build.toObject();
  }
}
