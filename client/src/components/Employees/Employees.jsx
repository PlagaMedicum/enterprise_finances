import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import DatePicker from "../DatePicker/DatePicker";
import Button from "react-bootstrap/Button";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";

class Employees extends React.Component {
  constructor(props){
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
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={3} pageSizeOptions={[3, 6]}/>
        </Jumbotron>
    );
  }
}

// TODO: Each row will be a <tr> with href'ed buttons to edit✎ and delete❌.

export default Employees;
