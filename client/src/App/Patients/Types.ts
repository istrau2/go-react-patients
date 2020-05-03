export interface Patient {
    MRN: Number,
    FirstName: string,
    LastName: string,
    DateOfBirth: Date
}

export const TOGGLE_LOADING = 'TOGGLE_LOADING'
export const RELOAD_RESULTS = 'RELOAD_RESULTS'

interface ToggleLoadingAction {
    type: typeof TOGGLE_LOADING,
    payload: {
        isLoading: boolean
    }
}

interface ReloadResultsAction {
    type: typeof RELOAD_RESULTS,
    payload: {
        count: Number,
        results: Patient[]
    }
}

export type PatientsActionTypes = ToggleLoadingAction | ReloadResultsAction