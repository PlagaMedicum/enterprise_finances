import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';
import Button from 'react-bootstrap/Button';
import ButtonToolbar from "react-bootstrap/ButtonToolbar";
import axios from 'axios';
import {ListGroup} from "react-bootstrap";
import ListGroupItem from "react-bootstrap/ListGroupItem";
import Spinner from "react-bootstrap/Spinner";

class RequestButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      error: null,
      isLoaded: false,
      items: []
    };
  }

  componentDidMount() {
    console.log("trying to post after mounting...");
    axios.post('http://localhost:1540/employee/add', {
      'name': 'Lavon',
      'position': 'headliner',
      'grade': 5,
      'tu-membership': false
    })
        .then(
            resp => {
              this.setState({
                isLoaded: true,
                items: resp.data
              });
              console.log(resp);
              console.log(this.state);
            })
        .catch(
            error => {
              this.setState({
                isLoaded: true,
                error
              });
              console.error(error);
            });
  }

  async handleClick(t) {
    await console.log("trying to post...");
    await axios.post('http://localhost:1540/employee/add', {
      'name': 'Lavon',
      'position': 'headliner',
      'grade': 5,
      'tu-membership': false
    })
        .then(
            resp => {
              t.setState({
                isLoaded: true,
                items: resp.data
              });
              console.log(resp);
              console.log(this.state);
            })
        .catch(
            error => {
              t.setState({
                isLoaded: true,
                error
              });
              console.error(error);
            });
  }


  render() {
    const {isLoaded, items} = this.state;
    if (!isLoaded) {
      return (
          <Spinner animation="border" variant="primary"/>
      );
    } else return (
        <ButtonToolbar>
          <Button variant="primary" size="lg" block onClick={() => this.handleClick(this)}>Go!</Button>
          <ListGroup>
            {Object.keys(items).map((key, index) => (
                <ListGroupItem key={key}>{key}. {items[key]}</ListGroupItem>
            ))}
          </ListGroup>
        </ButtonToolbar>
    );
  }
}

export default RequestButton;
