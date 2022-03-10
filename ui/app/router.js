import EmberRouter from '@ember/routing/router';
import config from 'hc-interviews/config/environment';

export default class Router extends EmberRouter {
  location = config.locationType;
  rootURL = config.rootURL;
}

Router.map(function () {
  this.route('interview', { path: '/' }, function () {
    this.route('begin', { path: '/interview/:interview_id' });
  });
});
