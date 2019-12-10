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
      items: [],
      add: false,
      edit: false,
      delete: false
    };
  }

  componentDidMount() {
    axios.get('http://localhost:1540/grade') // TODO: Add GET request for specific date
        .then(resp => this.setState({
          items: resp.data
        }))
        .catch(
            error => console.log(error)
        );
  }

  addElement() {
    // TODO: Showing modal and sending POST http request
    this.setState({add: true, edit: false, delete: false})
  }

  editElement() {
    // TODO: Showing modal and sending POST http request
    this.setState({add: false, edit: true, delete: false})
  }

  deleteElement() {
    // TODO: sending DELETE http request
    this.setState({add: false, edit: false, delete: true})
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
          <EditModal show={this.state.edit} hide={() => this.setState({edit: false})}/>
          <DeleteModal show={this.state.delete} hide={() => this.setState({delete: false})}/>
        </Jumbotron>
    );
  }
}

export default Grades;
