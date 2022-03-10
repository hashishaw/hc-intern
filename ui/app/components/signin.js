import Component from '@glimmer/component';
import { action } from '@ember/object';
import { service } from '@ember/service';
import { tracked } from '@glimmer/tracking';
export default class SigninComponent extends Component {
  @service router;

  @tracked code = '';

  @action
  handleSubmit(evt) {
    evt.preventDefault();
    if (!this.code) return;
    console.log(this.code, 'code');
    this.router.transitionTo('interview.begin', this.code);
  }
}
