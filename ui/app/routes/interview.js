import Route from '@ember/routing/route';

export default class InterviewRoute extends Route {
  model(params) {
    return fetch(`/api/interview/${params.id}`, {
      method: 'GET',
    }).then((res) => {
      console.log({ res });
      return res.json();
    });
  }
}
