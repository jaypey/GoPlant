import logo from './logo.svg';
import './App.css';
import NavbarLogged from './Navbar';
import {Container, Row, Col, Card} from 'react-bootstrap'
import { useEffect, useState } from 'react';
import SensorPickerCardList from './sensorPickerCardList';
import WithListLoading from './withListLoading';

function Body() {
    const ListLoading = WithListLoading(SensorPickerCardList);
    const [appState, setAppState] = useState({
        loading: false,
        sensors: null,
    });

    const getSensorsApiData = async() => {
        setAppState({loading:true});
        const response = await fetch(
            "http://localhost:8080/sensors"
        ).then((response) => response.json());
        console.log(response);
        setAppState({loading:false, sensors: response});
    };

    useEffect(() => {
        getSensorsApiData();
    },[setAppState]);

    return (
        <div className="Body">
        <NavbarLogged />
        <div className="Body-content">
            <Container>
                <ListLoading isLoading={appState.loading} sensors={appState.sensors} />
            </Container>
        </div>
        </div>
    );
}

export default Body;