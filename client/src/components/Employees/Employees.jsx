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
import DeleteModal from "../Modals/DeleteModal";

class Employees extends React.Component {

  // TODO:
  // 2. Add edit window with a PUT request inside
  // 4. Add addition window with a POST request inside
  // 5. Add GET request for specific date
  // 6. Add window of employee payments by dates

  constructor(props) {
    super(props);
    this.state = {
      host: 'http://localhost:1540',
      day: 0,
      month: 0,
      year: 0,
      items: [],
      adding: false,
      viewing: false,
      editing: false,
      deleting: false,
      id: null
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

  showView(id) {
    this.setState({adding: false, viewing: true, editing: false, deleting: false, id: id})
  }

  showEdit(id) {
    this.setState({adding: false, viewing: false, editing: true, deleting: false, id: id})
  }

  showDelete(id) {
    this.setState({adding: false, viewing: false, editing: false, deleting: true, id: id})
  }

  hide() {
    this.setState({adding: false, viewing: false, editing: false, deleting: false, id: null})
  }

  addElement(data) {
    axios.post(`${this.state.host}/employee/add`, data)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  deleteElement(id) {
    axios.delete(`${this.state.host}/employee/${id}/delete`)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
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
          <Button block variant="warning" onClick={() => this.showView(row.id)}>view</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="primary" onClick={() => this.showEdit(row.id)}>edit</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="danger" onClick={() => this.showDelete(row.id)}>delete</Button>
      )
    }];

    const dateSetters = {
      setDay: day => this.setState({day: day}),
      setMonth: month => this.setState({month: month}),
      setYear: year => this.setState({year: year})
    };

    return (
        <Jumbotron>
          <DatePicker setters={dateSetters}/>
          <Button block variant="success" onClick={() => this.showAdd()}>add</Button>
          <AddModal show={this.state.adding} hide={() => this.hide()} handler={data => this.addElement(data)}/>
          <p/>
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={10}
                      pageSizeOptions={[10, 20, 30]}/>
          <ViewModal id={this.state.id} show={this.state.viewing} hide={() => this.hide()}/>
          <EditModal show={this.state.editing} hide={() => this.hide()}/>
          <DeleteModal title="Delete Employee" show={this.state.deleting} hide={() => this.hide()}
                       confirm={() => this.deleteElement(this.state.id)}/>
        </Jumbotron>
    );
  }
}

export default Employees;
