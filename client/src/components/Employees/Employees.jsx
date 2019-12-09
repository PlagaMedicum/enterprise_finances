import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import DatePicker from "../DatePicker/DatePicker";
import Button from "react-bootstrap/Button";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";

class Employees extends React.Component {

  // TODO:
  // 1. Add edit✎ and delete❌ in each row(maybe in view window)
  // 2. Add edit window with a PUT request inside
  // 3. Add DELETE request
  // 4. Add addition window with a POST request inside
  // 5. Add GET request for specific date
  // 6. Add window of employee payments by dates

  constructor(props) {
    super(props);
    this.state = {
      items: []
    }
  }

  componentDidMount() {
    axios.get('http://localhost:1540/employee')
        .then(resp => this.setState({
          items: resp.data
        }))
        .catch(error => console.log(error))
  }

  render() {
    const data = this.state.items;
    const columns = [{
      Header: 'id',
      accessor: 'id'
    }, {
      Header: 'Name',
      accessor: 'name'
    }, {
      Header: 'Position',
      accessor: 'position'
    }, {
      Header: 'Salary',
      accessor: 'salary'
    }];

    return (
        <Jumbotron>
          <DatePicker/>
          <Button block variant="success">➕</Button>
          <p/>
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={3} pageSizeOptions={[10, 20, 30]}/>
        </Jumbotron>
    );
  }
}

export default Employees;
