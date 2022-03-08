import Component from '@glimmer/component';
import { action } from '@ember/object';
export default class InterviewPanelComponent extends Component {
  @action
  update(updates) {
    console.log({ updates });
  }
}
