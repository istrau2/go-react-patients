import React from "react";
import { NavLink } from "react-router-dom";
import { Navbar, Nav } from "react-bootstrap"
import { LinkContainer } from "react-router-bootstrap"

export default function NavBar () {
    return (
        <Navbar bg="light" expand="lg">
            <Navbar.Brand>
                ISSFST
            </Navbar.Brand>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
                <Nav className="mr-auto">
                    <LinkContainer exact={true} to="/">
                        <li className="nav-item">
                            <NavLink className="nav-link" exact={true} to="/">Patients</NavLink>
                        </li>
                    </LinkContainer>
                </Nav>
            </Navbar.Collapse>
        </Navbar>
    //     <nav className="navbar navbar-expand-lg navbar-light bg-light">
    //     <Link className="navbar-brand" to="/">TaskApp</Link>
    //     <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
    //       <span className="navbar-toggler-icon"></span>
    //     </button>
    //     <div className="collapse navbar-collapse" id="navbarNav">
    //       <ul className="navbar-nav">
    //         <LinkContainer exact={true} activeClassName='active' to="/">Patients</LinkContainer>
    //       </ul>
    //     </div>
    //   </nav>
    )
}