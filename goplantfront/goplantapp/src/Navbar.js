import Navbar from 'react-bootstrap/Navbar';
import Container from 'react-bootstrap/Container';
import {Nav, InputGroup, Button, Form} from 'react-bootstrap';

function NavbarTop(props) {
    return(
        <>
            <Navbar bg="dark" className="topBar">
                <Container>
                    <Navbar.Toggle aria-controls='basic-navbar-nav' />
                    <InputGroup className="w-25">
                        <Form.Control
                        placeholder="Search"
                        aria-label="Search"
                        aria-describedby="Search"
                        />
                        <Button variant="outline-secondary" id="button-addon2">
                        Search
                        </Button>
                    </InputGroup>
                    <Navbar.Collapse id="basic-navbar-nav" className="justify-content-end">
                        <Navbar.Text className='text-white'>
                            Signed in as: <a className='text-white' href="#login">JP Belval</a>
                        </Navbar.Text>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        </>
    );
}

export default NavbarTop;