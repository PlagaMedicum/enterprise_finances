import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import DatePicker from "../DatePicker/DatePicker";
import Button from "react-bootstrap/Button";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";
import AddModal from "./Modals/AddModal";
import ViewModal from "./Modals/ViewModal";
import EditModal from "./Modals/EditModal";
import DeleteModal from "./Modals/DeleteModal";

class Employees extends React.Component {

  // TODO:
  // 1. Add edit and delete in each row(maybe in view window)
  // 2. Add edit window with a PUT request inside
  // 3. Add DELETE request
  // 4. Add addition window with a POST request inside
  // 5. Add GET request for specific date
  // 6. Add window of employee payments by dates

  constructor(props) {
    super(props);
    this.state = {
      items: [],
      add: false,
      view: false,
      edit: false,
      delete: false
    }
  }

  componentDidMount() {
    axios.get('http://localhost:1540/employee') // TODO: Add GET request with specific date
        .then(resp => this.setState({
          items: resp.data
        }))
        .catch(error => console.log(error))
  }

  addElement() {
    this.setState({add: true, view: false, edit: false, delete: false})
  }

  viewElement() {
    this.setState({add: false, view: true, edit: false, delete: false})
  }

  editElement() {
    this.setState({add: false, view: false, edit: true, delete: false})
  }

  deleteElement() {
    this.setState({add: false, view: false, edit: false, delete: true})
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
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="warning" onClick={() => this.viewElement()}>view</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="primary" onClick={() => this.editElement()}>edit</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="danger" onClick={() => this.deleteElement()}>delete</Button>
      )
    }];

    return (
        <Jumbotron>
          <DatePicker/>
          <Button block variant="success" onClick={() => this.addElement()}>add</Button>
          <AddModal show={this.state.add} hide={() => this.setState({add: false})}/>
          <p/>
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={10}
                      pageSizeOptions={[10, 20, 30]}/>
          <ViewModal show={this.state.view} hide={() => this.setState({view: false})}/>
          <EditModal show={this.state.edit} hide={() => this.setState({edit: false})}/>
          <DeleteModal show={this.state.delete} hide={() => this.setState({delete: false})}/>
        </Jumbotron>
    );
  }
}

export default Employees;
