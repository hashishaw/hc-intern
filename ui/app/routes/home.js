import Route from '@ember/routing/route';

export default class HomeRoute extends Route {
  model() {
    return [
      { name: 'Chelsea Shaw', interview: 'intern-db' },
      { name: 'Penelope Luck', interview: 'kv-db' },
      { name: 'Eunice Foote', interview: 'kv-db' },
    ];
  }
}
