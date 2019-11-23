import React from "react";
import Table from "react-bootstrap/Table";
import Jumbotron from "react-bootstrap/Jumbotron";
import DatePicker from "../DatePicker/DatePicker";
import Button from "react-bootstrap/Button";

class Employees extends React.Component {
  render() {
    return (
        <Jumbotron>
          <DatePicker/>
          <Button block variant="success">➕</Button>
          <p/>
          <Table striped bordered hover variant="warning">
            <thead>
            <tr>
              <th>id</th>
              <th>Name</th>
              <th>Position</th>
              <th>Salary</th>
              <th/>
            </tr>
            </thead>
            <tbody>
            <tr>
              <th/>
            </tr>
            </tbody>
          </Table>
        </Jumbotron>
    );
  }
}

// TODO: Each row will be a <tr> with href'ed buttons to edit✎ and delete❌.

export default Employees;
