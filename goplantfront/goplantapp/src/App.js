import logo from './logo.svg';
import './App.css';
import NavbarLogged from './Navbar';
import Body from './Body';
import NavbarSide from './NavbarSide';
import {Container, Row, Col} from "react-bootstrap";
import "./Navbar.css";

function App() {
  return (
    <div className="App">
      <Container fluid className='h-100'>
        <Row className="h-100">
          <Col id="sidebar-wrapper">
            <NavbarSide className="bg-dark" />
          </Col>
          <Col id="page-content-wrapper">
            <Body/>
          </Col>
        </Row>
      </Container>
    </div>
  );
}

export default App;
