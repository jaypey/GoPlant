import './App.css';
import {Row, Col, Card} from 'react-bootstrap'
import { useEffect, useState } from 'react';

function SensorPickerCardList(props) {
    const {sensors} = props;
    
    if(sensors == null || sensors instanceof Error || sensors.length === 0) return <p>No sensors</p>;

    return (
        <Row className="cardPicker">
            {sensors.map((sensor) => {
                return(
                    <Col md={3}>
                        <Card key={sensor.ID} className="sensorCard mb-4" >
                            <Card.Body>
                                <Card.Title>{sensor.Name}</Card.Title>
                                <Card.Subtitle className="mb-2 text-muted">{sensor.IP}</Card.Subtitle>
                                <Card.Text>
                                Some quick example text to build on the card title and make up the
                                bulk of the card's content.
                                </Card.Text>
                                <Card.Link href="#">Card Link</Card.Link>
                                <Card.Link href="#">Another Link</Card.Link>
                            </Card.Body>
                        </Card>
                    </Col>
                );
            })}
    </Row>
    );
}

export default SensorPickerCardList;