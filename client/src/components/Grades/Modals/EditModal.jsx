import React, {useState} from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import DatePicker from "../../DatePicker/DatePicker";

function EditModal(props) {
  const style = {
    color: "black"
  };

  const [num, setNum] = useState(0);
  const [day, setDay] = useState(0);
  const [month, setMonth] = useState(0);
  const [year, setYear] = useState(0);
  const [coeff, setCoeff] = useState(0);

  const data = {
    'num': Number(num),
    'date': `${day}.${month}.${year}`,
    'coeff': Number(coeff)
  };

  // FIXME: Bug with updating default value after showing up

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>Edit Grade</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group controlId="formNum">
            <Form.Label>Order number</Form.Label>
            <Form.Control placeholder="grade" onChange={e => setNum(e.target.value)} defaultValue={props.data["num"]}/>
          </Form.Group>
          <Form.Group controlId="formDate">
            <Form.Label>Date</Form.Label>
            <DatePicker setters={{setDay, setMonth, setYear}}/>
          </Form.Group>
          <Form.Group controlId="formCoeff">
            <Form.Label>Coefficient</Form.Label>
            <Form.Control placeholder="coefficient" onChange={e => setCoeff(e.target.value)}
                          defaultValue={props.data["coeff"]}/>
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="danger" onClick={() => props.handler(data)}>Create</Button>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default EditModal;
