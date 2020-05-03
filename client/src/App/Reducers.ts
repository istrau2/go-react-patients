import {combineReducers} from 'redux';
import { PatientsReducer } from './Patients/Reducers'

export const RootReducer = combineReducers({
  patients: PatientsReducer
})

export type RootState = ReturnType<typeof RootReducer>