import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import Button from "react-bootstrap/Button";
import DatePicker from "../DatePicker/DatePicker";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";
import EditModal from "./Modals/EditModal";
import DeleteModal from "../Modals/DeleteModal";
import AddModal from "./Modals/AddModal";

class Grades extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      host: 'http://localhost:1540',
      day: 0,
      month: 0,
      year: 0,
      items: [],
      data: {},
      adding: false,
      editing: false,
      deleting: false,
      id: null
    };
  }

  componentDidMount() {
    this.updateTable()
  }

  updateTable() {
    axios.get(`${this.state.host}/grade`) // TODO: Add GET request for specific date
        .then(resp => {
          if (resp.data != null) {
            this.setState({
              items: resp.data
            })
          }
        })
        .catch(err => console.log(err));
  }

  addElement(data) {
    axios.post(`${this.state.host}/grade/add`, data)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  updateData(id) {
    axios.get(`${this.state.host}/grade/${id}`)
        .then(resp => {
          this.setState({data: resp.data});
          console.log(resp);
        })
        .catch(err => console.log(err))
  }

  editElement(data) {
    axios.post(`${this.state.host}/grade/${this.state.id}`, data)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  deleteElement(id) {
    axios.delete(`${this.state.host}/grade/${id}/delete`)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  showAdd() {
    this.setState({adding: true, editing: false, deleting: false})
  }

  showEdit(id) {
    this.updateData(id);
    this.setState({adding: false, editing: true, deleting: false, id: id})
  }

  showDelete(id) {
    this.setState({adding: false, editing: false, deleting: true, id: id})
  }

  hide() {
    this.setState({adding: false, editing: false, deleting: false, id: null})
  }

  render() {
    const data = this.state.items;
    const columns = [{
      Header: 'id',
      accessor: 'id'
    }, {
      Header: 'Grade',
      accessor: 'num'
    }, {
      Header: 'Date',
      accessor: 'date'
    }, {
      Header: 'Coefficient',
      accessor: 'coeff'
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
          <EditModal data={this.state.data} show={this.state.editing} hide={() => this.hide(this.state.id)}
                     handler={data => this.editElement(data)}/>
          <DeleteModal title="Delete Grade" show={this.state.deleting} hide={() => this.hide()}
                       confirm={() => this.deleteElement(this.state.id)}/>
        </Jumbotron>
    );
  }
}

export default Grades;
