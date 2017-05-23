import React, {Component} from 'react'
import { textarea, Field, FieldArray, reduxForm } from 'redux-form'
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import Select from 'react-select';
import { Link } from 'react-router';
import { HelpBlock } from 'react-bootstrap'
import 'react-select/dist/react-select.css';
import './styles.css';

class Incident_information extends Component {
	constructor(props) {
		super(props);
		this.state = {
		};
		this.showUpdateButton = this.showUpdateButton.bind(this);
	}

	componentDidMount() {
		this.props.actions.fetchIncidentInfo(this.props.params.incidentId)
	}

	showUpdateButton() {
		if(this.props.state.incidents.Incidents.Status == "active") {
			return(	<div>
								<form onSubmit={ this.handleSubmit }>
									<Link to={`incidents/${this.props.params.incidentId}/update`} > <button className="btn btn-info floatRight">Add Update</button> </Link>
								</form>
							</div>
						);
		}
	}

	render() {
		//======================== style =============================================
		let borderStyle1 = {
			border: 'solid',
			borderWidth: '2px',
			padding: '20px 25px 0px 100px'
		};

		let borderStyle = {
			border: 'solid',
			borderWidth: '1px',
			padding: '20px 25px 0px 100px'
		};
		//============================================================================

		let Recorder = '';
		let Status = 'Resolved';
		let Description = '';
		let Component = '';
		let IncidentUpdates = [];
		if(this.props.state.incidents.Incidents.Component) {
			Recorder = this.props.state.incidents.Incidents.Recorder
			Description = this.props.state.incidents.Incidents.Description
			Component = this.props.state.incidents.Incidents.Component
			if(this.props.state.incidents.Incidents.Status == 'active') {
				Status = "Active"
			}
			IncidentUpdates = this.props.state.incidents.Incidents.IncidentUpdates
		}

		const { handleSubmit, pristine, reset, submitting } = this.props
		return (

			<div>
				{this.showUpdateButton()}
					<div>
					<h3 className="center"><u> Incident #{ this.props.params.incidentId} </u></h3>
					<div style={borderStyle1}>
						<div className="clearfix">
							<div className = "col-lg-2 col-lg-offset-1">
								<label>Status :</label>
							</div>
							<div className = "col-lg-2 fontStyle">
								{ Status }
							</div>
						</div>
						<div className="clearfix">
							<div className="col-lg-2 col-lg-offset-1">
								<label >Recorder :</label>
							</div>
							<div className = "col-lg-2 fontStyle" >
								{ Recorder }
							</div>
						</div>
						<div className="clearfix">
							<div className="col-lg-2 col-lg-offset-1">
								<label >Component :</label>
							</div>
							<div className = "col-lg-2 fontStyle">
								{ Component }
							</div>
						</div>
						<div className="clearfix form-group">
							<div className="col-lg-2 col-lg-offset-1">
								<label >Description :</label>
							</div>
							<div className="col-lg-8 fontStyle">
								{ Description }
							</div>
						</div>
					</div>
				</div>

				<h3 className="center"> Updates </h3>

			{/*================== Display incident updates ...=====================*/}

				{IncidentUpdates.map(t => (
					<div key={t.Description}>
						<br/>
						<div> Date : {t.Resolved_Date}</div>
						<div style={borderStyle}>
							<div className="clearfix">
								<div className = "col-lg-2 col-lg-offset-1">
									<label >Update By :</label>
								</div>
								<div className = "col-lg-2 fontStyle">
									{ t.Updated_by }
								</div>
							</div>
							<div className="clearfix form-group">
								<div className = "col-lg-2 col-lg-offset-1">
									<label >Description :</label>
								</div>
								<div className = "col-lg-8 fontStyle">
									<textarea
										className="textAreaWidth"
										readOnly={ true }
										value={t.Description}>
									</textarea>
								</div>
							</div>
						</div>
					</div>
				))}

			{/*====================================================================*/}

			</div>
		)
	}
}

export default reduxForm({
	form: 'Incident_information',// a unique identifier for this form
})(Incident_information)

