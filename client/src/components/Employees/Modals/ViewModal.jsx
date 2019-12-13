import React, {useState} from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";
import ReactTable from "react-table";
import "react-table/react-table.css";
import DatePicker from "../../DatePicker/DatePicker";
import timeConverter from "../../../timeConverter";

function ViewModal(props) {
  const style = {
    color: "black"
  };

  const columns = [{
    Header: 'Date',
    accessor: 'date',
  }, {
    Header: 'Accruals',
    accessor: 'accruals',
  }, {
    Header: 'Deduction',
    accessor: 'deduction',
  }, {
    Header: 'Salary',
    accessor: 'salary',
  }];

  const [day, setDay] = useState(1);
  const [month, setMonth] = useState(1);
  const [year, setYear] = useState(1);

  const data = {
    'id': props.data["id"],
    'date': timeConverter(day, month, year),
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>View Employee</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <b>#{props.data["id"]}</b><br/>
          <b>Name:</b> {props.data["name"]}<br/>
          <b>Position:</b> {props.data["position"]}<br/>
          <b>Grade:</b> {props.data["grade"]}<br/>
          <b>{props.data["tu-membership"] ? "Is" : "Is not"} Trade Union member</b><br/>
          <b>Salary:</b><br/>
          <DatePicker setters={{setDay, setMonth, setYear}}/>
          <Button block variant="secondary" onClick={() => props.update(data)}>update</Button><br/>
          <ReactTable style={style} data={props.data["payments"]} columns={columns} defaultPageSize={5}
                      pageSizeOptions={[5, 10, 20, 30]}/>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default ViewModal;
