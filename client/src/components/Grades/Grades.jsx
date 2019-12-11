import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import Button from "react-bootstrap/Button";
import DatePicker from "../DatePicker/DatePicker";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";
import EditModal from "./Modals/EditModal";
import DeleteModal from "./Modals/DeleteModal";
import AddModal from "./Modals/AddModal";

class Grades extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      host: 'http://localhost:1540',
      items: [],
      adding: false,
      editing: false,
      deleting: false,
      row: null
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

  showAdd() {
    // TODO: Showing modal and sending POST http request
    this.setState({adding: true, editing: false, deleting: false})
  }

  showEdit() {
    // TODO: Showing modal and sending POST http request
    this.setState({adding: false, editing: true, deleting: false})
  }

  showDelete(row) {
    // TODO: sending DELETE http request
    this.setState({adding: false, editing: false, deleting: true, row: row})
  }

  deleteElement(row) {
    axios.delete(`${this.state.host}/grade/${row.id}/delete`)
        .then(resp => console.log(resp))
        .catch(err => console.log(err));
    this.setState({deleting: false, row: null})
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
          <EditModal show={this.state.editing} hide={() => this.setState({editing: false})}/>
          <DeleteModal show={this.state.deleting} hide={() => this.setState({deleting: false, row: null})}
                       confirm={() => this.deleteElement(this.state.row)}/>
        </Jumbotron>
    );
  }
}

export default Grades;
