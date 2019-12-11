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
      host: 'http://localhost:1540',
      items: [],
      adding: false,
      viewing: false,
      editing: false,
      deleting: false,
      row: null
    }
  }

  componentDidMount() {
    this.updateTable()
  }

  updateTable() {
    axios.get(`${this.state.host}/employee`) // TODO: Add GET request with specific date
        .then(resp => {
          if (resp.data != null) {
            this.setState({
              items: resp.data
            })
          }
        })
        .catch(err => console.log(err))
  }

  showAdd() {
    this.setState({adding: true, viewing: false, editing: false, deleting: false})
  }

  showView() {
    this.setState({adding: false, viewing: true, editing: false, deleting: false})
  }

  showEdit() {
    this.setState({adding: false, viewing: false, editing: true, deleting: false})
  }

  showDelete(row) {
    this.setState({adding: false, viewing: false, editing: false, deleting: true, row: row})
  }

  deleteElement(row) {
    axios.delete(`${this.state.host}/employee/${row.id}/delete`)
        .then(resp => console.log(resp))
        .catch(err => console.log(err));
    this.setState({deleting: false, row: null});
    this.updateTable();
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
          <Button block variant="warning" onClick={() => this.showView()}>view</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="primary" onClick={() => this.showEdit()}>edit</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="danger" onClick={() => this.showDelete(row)}>delete</Button>
      )
    }];

    return (
        <Jumbotron>
          <DatePicker/>
          <Button block variant="success" onClick={() => this.showAdd()}>add</Button>
          <AddModal show={this.state.adding} hide={() => this.setState({adding: false})}/>
          <p/>
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={10}
                      pageSizeOptions={[10, 20, 30]}/>
          <ViewModal show={this.state.viewing} hide={() => this.setState({viewing: false})}/>
          <EditModal show={this.state.editing} hide={() => this.setState({editing: false})}/>
          <DeleteModal show={this.state.deleting} hide={() => this.setState({deleting: false, row: null})}
                       confirm={() => this.deleteElement(this.state.row)}/>
        </Jumbotron>
    );
  }
}

export default Employees;
