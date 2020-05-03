import { TOGGLE_LOADING, RELOAD_RESULTS, Patient, PatientsActionTypes } from './Types';

export interface PatientsState {
    isLoading: boolean,
    count: Number,
    results: Patient[]
}

const initialPatientsState: PatientsState = {
    isLoading: false,
    count: 0,
    results: []
}

export function PatientsReducer(
    state: PatientsState = initialPatientsState,
    action: PatientsActionTypes
) : PatientsState {
    switch (action.type) {
        case TOGGLE_LOADING:
            case RELOAD_RESULTS:
            return {
                ...state,
                ...action.payload
            }
        default:
            return state
    }
}