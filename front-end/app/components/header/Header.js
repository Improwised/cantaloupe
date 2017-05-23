import React, { Component } from 'react';
import { Link } from 'react-router';
import './styles.css';

class Header extends Component {

	render() {
		return (
			<div className="container-fluid">
				<nav className="navbar navbar-inverse">
					<div className="container">

						<div className="navbar-header">
							<button type="button" className="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar-collapse-3">
								<span className="sr-only">Toggle navigation</span>
								<span className="icon-bar"></span>
								<span className="icon-bar"></span>
								<span className="icon-bar"></span>
							</button>
							<a className="navbar-brand" href="#">Asset Tracking System</a>
						</div>
						<div className="collapse navbar-collapse" id="navbar-collapse-3">
							<ul className="nav navbar-nav navbar-right">
								<li><Link to={`/category`}> Categories </Link></li>
								<li><Link to={`/machines`}> Machines </Link></li>
								<li><Link to={`/users`}> Users </Link></li>
								<li><Link to={`/invoices`}> Invoices </Link></li>
								<li><Link to={`/components`}> Components </Link></li>
								<li><Link to={`/incidents`}> Incidents </Link></li>
							</ul>
						</div>
					</div>
				</nav>
			</div>
		)
	}
}

export default Header
