import React, { Component } from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Alert from 'react-bootstrap/Alert';
import RequestButton from './uButton';

class App extends Component {
  render() {
    return (
      <div className="App">
        <Alert variant="success">
          <Alert.Heading>Ah, ello there, Gordn Freeman!</Alert.Heading>
          <p>
            Why do we all must to wear those ridiculous ties?!
          </p>
          <hr />
          <p className="mb-0">
            Did you see my coffeecup?
          </p>
        </Alert>
        <RequestButton />
      </div>
    );
  }
}

export default App;
