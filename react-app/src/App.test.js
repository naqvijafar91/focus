import React from 'react';
import ReactDOM from 'react-dom';
import UserStore from './LoginPage/userStore';
import App from './App';
UserStore.saveUser({token:"sample-test"});

it('renders without crashing', () => {
  const div = document.createElement('div');
  ReactDOM.render(<App />, div);
  ReactDOM.unmountComponentAtNode(div);
});

it('should parse server response',()=>{
  let resp = {
    "data": [
      {
        "id": "5f9be853-a86c-478e-b937-f5c4a14a8f43",
        "name": "Work",
        "remaining_tasks": 11,
        "tasks": [
          {
            "id": "c50e63db-cf3a-49da-a145-d986a32a6d3b",
            "description": "Now Updating",
            "folder_id": "27f526e6-8478-4806-8bdc-08e8d4e428be",
            "due_date": "01-05-2020",
            "completed_date": "0001-01-01T00:00:00Z"
          }
        ]
      }
    ]
  };
  const testApp = new App();
  const parsed = testApp.parseCompleteServerResponse(resp);
  const parsedDate = parsed.data[0].tasks[0].due_date;
  expect(parsedDate.getMonth()).toBe(4);
  expect(parsedDate.getDate()).toBe(1);
  expect(parsedDate.getFullYear()).toBe(2020);

});