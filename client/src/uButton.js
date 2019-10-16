import React from 'react';
import Button from 'react-bootstrap/Button';
import axios from "axios"

const instance = axios.create({
    proxy: {
        host: '127.0.0.1',
        port: 1540,
    },
    timeout: 1000,
});

class uButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = {isButtonOn: true};

    // This binding is necessary to make `this` work in the callback
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick() {
      instance.get('/')
      .then(function (response) {
        this.setState(
          state => ({
            isButtonOn: !state.isButtonOn,
          }));
        console.log(response);
      })
      .catch(function (error) {
        console.log(error);
      });
  }

  render() {
    return (
      <Button
        variant="primary"
        size="lg"
        block
        onClick={this.handleClick}>
        {this.state.isButtonOn ? 'ON' : 'OFF'}
      </Button>
    );
  }
}

export default uButton;
