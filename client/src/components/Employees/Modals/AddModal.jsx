import React, {useState} from "react";
import Modal from "react-bootstrap/Modal";
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form"

function AddModal(props) {
  const style = {
    color: "black"
  };

  const [name, setName] = useState("");
  const [position, setPosition] = useState("");
  const [grade, setGrade] = useState("");
  const [tum, setTum] = useState(false);

  const data = {
    'name': name,
    'position': position,
    'grade': Number(grade),
    'tu-membership': tum
  };

  return (
      <Modal style={style} show={props.show} onHide={props.hide}>
        <Modal.Header>
          <Modal.Title>Add Employee</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group controlId="formName">
            <Form.Label>Name</Form.Label>
            <Form.Control placeholder="name" onChange={e => setName(e.target.value)}/>
          </Form.Group>
          <Form.Group controlId="formPosition">
            <Form.Label>Position</Form.Label>
            <Form.Control placeholder="position" onChange={e => setPosition(e.target.value)}/>
          </Form.Group>
          <Form.Group controlId="formGrade">
            <Form.Label>Grade</Form.Label>
            <Form.Control placeholder="grade" onChange={e => setGrade(e.target.value)}/>
          </Form.Group>
          <Form.Group controlId="fromTum">
            <Form.Check type="checkbox" label="Trade Union Membership" onChange={e => setTum(!tum)}/>
          </Form.Group>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="success" onClick={() => props.handler(data)}>Create</Button>
          <Button variant="danger" onClick={props.hide}>Close</Button>
        </Modal.Footer>
      </Modal>
  );
}

export default AddModal;
