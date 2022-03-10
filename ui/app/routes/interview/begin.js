import Route from '@ember/routing/route';

export default class InterviewBeginRoute extends Route {
  model(params) {
    return fetch(`/api/interview/${params.interview_id}`, {
      method: 'GET',
    }).then((res) => {
      console.log({ res });
      return res.json();
    });
  }
}
