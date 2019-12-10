import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import Button from "react-bootstrap/Button";
import DatePicker from "../DatePicker/DatePicker";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";
import EditModal from "./EditModal";

class Grades extends React.Component {
  constructor(props){
    super(props);
    this.state = {
      items: [],
      edit: false
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

  addElement(e) {
    // TODO: Showing modal and sending POST http request
  }

  editElement(e, row) {
    // TODO: Showing modal and sending POST http request
    this.setState({edit: true})
  }

  deleteElement(e, row) {
    // TODO: Showing modal and sending DELETE http request
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
          <Button block variant="primary" onClick={e => this.editElement(e, row)}>edit</Button>
      )
    }, {
      Header: "",
      width: 100,
      Cell: ({row}) => (
          <Button block variant="danger" onClick={e => this.deleteElement(e, row)}>delete</Button>
      )
    }];

    return (
        <Jumbotron>
          <DatePicker/>
          <Button block variant="success" onClick={e => this.addElement(e)}>add</Button>
          <p/>
          <ReactTable style={{color: 'black'}} data={data} columns={columns} defaultPageSize={10}
                      pageSizeOptions={[10, 20, 30]}/>
          <EditModal show={this.state.edit} hide={() => {
            this.setState({show: false})
          }}/>
        </Jumbotron>
    );
  }
}

export default Grades;
