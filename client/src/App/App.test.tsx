import React from 'react';
import {Provider} from 'react-redux';
import { render } from '@testing-library/react';
import App from './App';
import {createStore} from 'redux';
import {RootReducer} from './Reducers';

const store = createStore(RootReducer)

test('renders learn react link', () => {
  const { getByText } = render(<Provider store={store}><App /></Provider>);
  const linkElement = getByText(/Patients/i);
  expect(linkElement).toBeInTheDocument();
});
