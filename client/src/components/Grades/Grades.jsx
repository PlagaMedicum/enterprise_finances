import React from "react";
import Table from "react-bootstrap/Table";
import Jumbotron from "react-bootstrap/Jumbotron";
import Button from "react-bootstrap/Button";
import DatePicker from "../DatePicker/DatePicker";

class Grades extends React.Component {
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
              <th>Grade</th>
              <th>Coefficient</th>
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

export default Grades;
