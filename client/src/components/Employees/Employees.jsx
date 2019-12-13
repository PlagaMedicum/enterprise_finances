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
import timeConverter from "../../timeConverter";

class Employees extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      host: 'http://localhost:1540',
      day: 1,
      month: 1,
      year: 1,
      items: [],
      data: [],
      adding: false,
      viewing: false,
      editing: false,
      deleting: false,
    }
  }

  hide() {
    this.setState({adding: false, viewing: false, editing: false, deleting: false})
  }

  updateTable() {
    axios.get(`${this.state.host}/employee?date=${timeConverter(this.state.day, this.state.month, this.state.year)}`)
        .then(resp => {
          if (resp.data != null) {
            this.setState({
              items: resp.data
            })
          }
        })
        .catch(err => console.log(err))
  }

  getEmployee(id) {
    axios.get(`${this.state.host}/employee/${id}`)
        .then(resp => {
          this.setState({data: resp.data});
          console.log(resp);
        })
        .catch(err => console.log(err))
  }

  getPayments(data) {
    let r = this.state.data;
    axios.get(`${this.state.host}/employee/${data["id"]}/payments?date=${data["date"]}`)
        .then(resp => {
          if (resp.data != null) {
            r["payments"] = resp.data;
          } else {
            r["payments"] = [];
          }
          this.setState({data: r});
          console.log(resp);
        })
        .catch(err => console.log(err))
  }

  editEmployee(data) {
    axios.post(`${this.state.host}/employee/${this.state.data["id"]}`, data)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  addEmployee(data) {
    axios.post(`${this.state.host}/employee/add`, data)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  deleteEmployee(id) {
    axios.delete(`${this.state.host}/employee/${id}/delete`)
        .then(resp => {
          this.updateTable();
          console.log(resp);
        })
        .catch(err => console.log(err));
    this.hide();
  }

  showAdd() {
    this.hide();
    this.setState({adding: true});
  }

  showView(id) {
    this.getEmployee(id);
    this.hide();
    this.setState({viewing: true, data: {"id": id}});
  }

  showEdit(id) {
    this.getEmployee(id);
    this.hide();
    this.setState({editing: true, data: {"id": id}})
  }

  showDelete(id) {
    this.hide();
    this.setState({deleting: true, data: {"id": id}})
  }

  componentDidMount() {
    this.updateTable()
  }

  render() {
    const style = {
      color: "black"
    };

    const data = this.state.items;
    const columns = [{
      Header: 'id',
      accessor: 'id'
    }, {
      Header: 'Date',
      accessor: 'date'
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
      setYear: year => this.setState({year: year}),
    };

    return (
        <Jumbotron>
          <DatePicker setters={dateSetters}/>
          <Button block variant="secondary" onClick={() => this.updateTable()}>update</Button><br/>
          <Button block variant="success" onClick={() => this.showAdd()}>add</Button>
          <AddModal show={this.state.adding} hide={() => this.hide()} handler={data => this.addEmployee(data)}/>
          <ReactTable style={style} data={data} columns={columns} defaultPageSize={10}
                      pageSizeOptions={[10, 20, 30]}/>
          <ViewModal data={this.state.data} show={this.state.viewing} hide={() => this.hide()}
                     update={data => this.getPayments(data)}/>
          <EditModal data={this.state.data} show={this.state.editing} hide={() => this.hide()}
                     handler={data => this.editEmployee(data)}/>
          <DeleteModal title="Delete Employee" show={this.state.deleting} hide={() => this.hide()}
                       confirm={() => this.deleteEmployee(this.state.data["id"])}/>
        </Jumbotron>
    );
  }
}

export default Employees;
