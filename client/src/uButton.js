import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import ButtonToolbar from "react-bootstrap/ButtonToolbar";
import axios from "axios"

class RequestButton extends React.Component {
  async handleClick() {
    axios.get('http://localhost:1540/')
      .then(function (response) {
        console.log(response);
      })
      .catch(function (error) {
        console.error(error);
      });
  }

  render() {
    return (
      <ButtonToolbar>
        <Button
          variant="primary"
          size="lg"
          block
          onClick={this.handleClick}>
          Call
        </Button>
      </ButtonToolbar>
    );
  }
}

export default RequestButton;
