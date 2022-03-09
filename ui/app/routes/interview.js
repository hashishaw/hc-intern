import Route from '@ember/routing/route';

export default class InterviewRoute extends Route {
  model(params) {
    console.log({ params });
    // throw new Error('Whoops bad');
    // return fetch(`/api/candidate/${params.id}`);
    return {
      name: 'Chelsea Shaw',
      code: params.id,
    };
  }
}
