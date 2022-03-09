import Component from '@glimmer/component';
import { action } from '@ember/object';
import { service } from '@ember/service';

export default class SigninComponent extends Component {
  @service router;

  @action
  handleSubmit() {
    console.log(this.code, 'code');
    this.router.transitionTo('interview', this.code);
  }
}
