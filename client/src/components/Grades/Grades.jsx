import React from "react";
import Jumbotron from "react-bootstrap/Jumbotron";
import Button from "react-bootstrap/Button";
import DatePicker from "../DatePicker/DatePicker";
import axios from "axios";
import ReactTable from "react-table";
import "react-table/react-table.css";

class Grades extends React.Component {

  // TODO:
  // 1. Add edit✎ and delete❌ in each row
  // 2. Add edit window with a PUT request inside
  // 3. Add DELETE request
  // 4. Add addition window with a POST request inside
  // 5. Add GET request for specific date

  constructor(props){
    super(props);
    this.state = {
      items: []
    };
  }

  componentDidMount() {
    axios.get('http://localhost:1540/grade')
        .then(resp => this.setState({
                items: resp.data
              }))
        .catch(
            error => console.log(error)
        );
  }

  render() {
    const data = this.state.items;
    const columns = [{
      Header: 'id',
      accessor: 'id'
    }, {
      Header: 'Grade',
      accessor: 'grade'
    }, {
      Header: 'Coefficient',
      accessor: 'coefficient'
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

export default Grades;
