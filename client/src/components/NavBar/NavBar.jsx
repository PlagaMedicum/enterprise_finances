import React from "react";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";

class NavBar extends React.Component {
  render() {
    return (
        <Navbar bg="warning" variant="light">
          <Navbar.Brand href="/">Black Mesa</Navbar.Brand>
          <Nav className="mr-auto">
            <Nav.Link href="/employees">Employees</Nav.Link>
            <Nav.Link href="/grades">Grades</Nav.Link>
          </Nav>
        </Navbar>
    );
  }
}

export default NavBar;
