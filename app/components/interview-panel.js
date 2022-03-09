import Component from '@glimmer/component';
import { action } from '@ember/object';
import CodeMirror from 'codemirror';
export default class InterviewPanelComponent extends Component {
  get partOne() {
    return `Set(name, value) {}
    Get(name) {}
    Unset(name) {
      return 0;
    }
    `;
  }

  get document() {
    return CodeMirror.Doc(this.partOne, { name: 'javascript', json: true });
  }
  @action
  update(updates) {
    console.log({ updates });
  }
}
