import Service from '@ember/service';
import { tracked } from '@glimmer/tracking';

export default class ThemeService extends Service {
  @tracked isDark = false;

  toggleTheme() {
    this.isDark = !this.isDark;
  }
}
