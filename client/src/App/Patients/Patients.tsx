import 'react-bootstrap-typeahead/css/Typeahead.css';
import React from "react";
import moment from "moment";
import { useSelector, useDispatch } from 'react-redux'
import { Container} from "react-bootstrap"
import {AsyncTypeahead, Highlighter } from 'react-bootstrap-typeahead';
import {TOGGLE_LOADING, RELOAD_RESULTS, Patient} from './Types';
import {RootState} from '../Reducers';
import {BACKEND} from '../Constants'

export default function Patients() {
    const state = useSelector((s : RootState) => s.patients)
    const dispatch = useDispatch()
    
    return (
        <Container className="p-3">
            <div className="mx-auto" style={{width: "100%", maxWidth: "500px"}}>
                <AsyncTypeahead id="patients_typeahead" 
                    isLoading={state.isLoading}
                    labelKey={option => option.FirstName + option.LastName}
                    placeholder="Search for a patient..."
                    onSearch={query => {
                        dispatch({ type: TOGGLE_LOADING, payload: { isLoading: true }});
                        fetch(`${BACKEND}/v1/patient?match=${query}`)
                        .then(resp => resp.json())
                        .then(json => {
                            dispatch({ type: RELOAD_RESULTS, payload: { count: json.Count, results: json.List }});
                            dispatch({ type: TOGGLE_LOADING, payload: { isLoading: false }});
                        });
                    }}
                    renderMenuItemChildren={renderMenuItemChildren}
                    options={state.results}
                />
            </div>
        </Container>
    );
  }

function renderMenuItemChildren(option : Patient, props : any, index: Number) {
    return [
        <Highlighter key="name" search={props.text}>
          {firstCharUpper(option.FirstName) + " " + firstCharUpper(option.LastName)}
        </Highlighter>,
        <span key="mrn" className="float-right">
            <small>
                <span className="text-muted">MRN</span>&nbsp;<strong>{option.MRN}</strong>
            </small>
        </span>,
        <div key="dob">
          <small>
            {moment(option.DateOfBirth).format("MM/DD/YYYY")}
          </small>
        </div>,
      ];
}

function firstCharUpper(str : string) : string {
    return str.slice(0, 1).toUpperCase() + str.slice(1);
}  