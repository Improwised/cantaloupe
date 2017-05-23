import React, {Component} from 'react'
import { textarea, Field, FieldArray, reduxForm } from 'redux-form'
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import Select from 'react-select';
import DatePicker from 'react-datepicker';
import moment from 'moment';
import { HelpBlock } from 'react-bootstrap';

import 'react-datepicker/dist/react-datepicker.css';
import 'react-select/dist/react-select.css';
import './styles.css';

class Edit_invoice extends Component {
	constructor(props) {
		super(props);
		this.state = {
			inputs :[],
			date: moment(),

			invoiceErr:'',
			invoicerErr:'',
			contactErr:'',
			serialErr:'',
			componentErr:'',
			category:''
		};
		this.handleFields = this.handleFields.bind(this);
		this.handleSubmit = this.handleSubmit.bind(this);
		this.handleInvoiceDateChange = this.handleInvoiceDateChange.bind(this);
		this.addInputField = this.addInputField.bind(this);
		this.removeInputField = this.removeInputField.bind(this);
		this.handleCategoryChange = this.handleCategoryChange.bind(this);
		this.handleDateChange = this.handleDateChange.bind(this);
		this.handleFieldsComponent = this.handleFieldsComponent.bind(this);
	}

	handleInvoiceDateChange(date) {
		this.setState({date: date});
		let id = `Invoice_date`
		this.props.actions.setFieldForInvoice(id, date);
		this.props.actions.fetchCategories(false, this.props.dispatch);
	}

	handleFieldsComponent(event) {
		this.props.actions.setFieldForComponent(event.target.id, event.target.value);
	}

	handleFields(event) {
		this.props.actions.setFieldForInvoice(event.target.id, event.target.value);
	}

	handleDateChange(componentId, date) {
		this.setState({startDate: date});
		let id = `date_${ componentId }`
		this.props.actions.setFieldForComponent(id, date);

	}

	handleCategoryChange(componentId, event) {
		this.setState({category: event.label});
		let id = `category_${ componentId }`
		this.props.actions.setFieldForComponent(id, event.value);
	}

	addInputField(e) {
		e.preventDefault();
		var inputs = this.state.inputs;
		inputs.push({name: null});
		this.setState({inputs : inputs});
	}

	removeInputField(index) {
		var inputs = this.state.inputs;
		inputs.splice(index, 1);
		this.setState({inputs : inputs});
	}

	handleSubmit(e) {
		e.preventDefault(e);
		this.setState({invoiceErr: ''});
		this.setState({invoicerErr: ''});
		this.setState({contactErr: ''});
		this.setState({serialErr: ''});
		this.setState({componentErr: ''});

		let status = true;
		let data = this.props.props.invoices
		if(data.invoice == '') {
			this.setState({invoiceErr: '*Required'});
			status = false;
		}
		if(data.invoicer == '') {
			this.setState({invoicerErr: '*Required'})
			status = false;
		}
		if(data.contact == '' || data.contact.length != 10 || isNaN(data.contact)) {
			this.setState({contactErr: '*Enter Valid Contact'})
			status = false;
		}
		if (status) {
			this.props.actions.editInvoice(this.props.params.invoiceId, this.props.props.invoices)
		}
	}

	componentDidMount() {
		let { invoiceId } = this.props.params;
		this.props.actions.fetchInvoiceDetails(invoiceId);
		this.props.actions.fetchCategories(false, this.props.dispatch);
	}

	render() {
		//============style=========================================================

		let borderStyle = {
			border: 'solid',
			borderWidth: '2px',
			padding: '20px 25px 20px 100px'
		};

		let selectcss ={
			width: '80%',
			overflow: 'scroll'
		}

		//==========================================================================

		var invoiceDate = moment(this.props.props.invoices.date);
		let invoice;
		let invoicer;
		let contact;
		let description;
		let address;
		let components = [];

		let allInvoiceData = this.props.props.invoices.data;

		let categories = [];
		for(let i = 0; i < this.props.props.category.AllCategories.length; i++){
			let ComponentInfo = { value : this.props.props.category.AllCategories[i].Id, label: this.props.props.category.AllCategories[i].Category };
			categories[i] = ComponentInfo;
		}

		if(this.props.props.invoices.invoice || this.props.props.invoices.invoicer) {
			invoice = this.props.props.invoices.invoice
			invoicer = this.props.props.invoices.invoicer
			contact = this.props.props.invoices.contact
			address = this.props.props.invoices.address
			description = this.props.props.invoices.description
			components = this.props.props.invoices.components
			invoiceDate = moment(this.props.props.invoices.Invoice_date)
		}

		const { handleSubmit, pristine, reset, submitting } = this.props
		return (
			<div >
				<form onSubmit={ this.handleSubmit }>
					<button className="btn btn-info floatRight" >Save</button>
				</form>
				<button className="btn btn-success pull-right fixPositionForAdd" onClick={this.addInputField}> Add Component </button>

				<div>
				<h2 className="center"> Edit Invoice </h2>
				<div style={borderStyle}>
					<div className="clearfix form-group">
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Invoice*</label>

							<input
								className="textboxSize"
								type="text"
								name="Invoice"
								id="invoice"
								value={ invoice || '' }
								onChange={this.handleFields}
								placeholder="ex. 12MOUSE1811"
							/>

							<HelpBlock className="errFontStyle"> {this.state.invoiceErr} </HelpBlock>
						</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Invoicer*</label>

							<input
								className="textboxSize"
								type="text"
								name="Invoicer"
								id="invoicer"
								value={ invoicer || '' }
								onChange={this.handleFields}
								placeholder="ex. Jay systems"
							/>

							<HelpBlock className="errFontStyle"> {this.state.invoicerErr} </HelpBlock>
						</div>
					</div>
					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Invoice Date*</label>

							<DatePicker
								className="textboxSize"
								name="Invoice_date"
								id="invoice_date"
								onChange={this.handleInvoiceDateChange}
								selected= { invoiceDate }
								peekNextMonth
								showYearDropdown
								showMonthDropdown
								dropdownMode="select"
								dateFormat="DD/MM/YYYY"
						 	/>

						</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Contact*</label>

							<input
								className="textboxSize"
								type="text"
								name="Contact"
								id="contact"
								value={ contact || '' }
								onChange={this.handleFields}
								placeholder="ex. +91 9909970574"
							/>

							<HelpBlock className="errFontStyle"> {this.state.contactErr} </HelpBlock>
						</div>
					</div>
					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Description</label>

							<textarea
								className="textAreaSize"
								name="Description"
								id="description"
								value={ description || '' }
								onChange={this.handleFields}
								placeholder="description"
							/>

						</div>
						<div className="col-lg-2 col-lg-offset-2">
							<label >Address</label>

							<textarea
								className="textAreaSize"
								name="Address"
								id="address"
								value={ address || '' }
								onChange={this.handleFields}
								placeholder="address"
							/>

						</div>
					</div>
				</div>
			</div>

				<div>
				<div className="center">
					<h3> Component(s) </h3>
				</div>
				<div>
					<BootstrapTable
						data={components}
						pagination={true}
						striped={true}
						search={true}
						hover={true}>
						<TableHeaderColumn isKey={true} hidden={true} dataSort={true} dataField="Id">#</TableHeaderColumn>
						<TableHeaderColumn width="170" dataSort={true} dataField="Serial_no">Serial</TableHeaderColumn>
						<TableHeaderColumn width="150" dataSort={true} dataField="Name">Name</TableHeaderColumn>
						<TableHeaderColumn width="80" dataSort={true} dataField="Active">Active</TableHeaderColumn>
						<TableHeaderColumn width="100" dataField="Warranty_till">Warranty Till</TableHeaderColumn>
						<TableHeaderColumn hidden={true} dataField="Invoice_id">Invoice</TableHeaderColumn>
						<TableHeaderColumn width="280" dataSort={true} dataField="Description">Description</TableHeaderColumn>
						<TableHeaderColumn width="100" dataSort={true} dataField="Category">Category</TableHeaderColumn>
						<TableHeaderColumn width="100" dataSort={true} dataField="Machine">Machine</TableHeaderColumn>
					</BootstrapTable>
				</div>

				<div className="setWidth">
					<label className="control-label text-center">Components </label>
					<br/>
					<HelpBlock className="center errFontStyle"> { this.state.componentErr } </HelpBlock>
					<br/>

					{this.state.inputs.map(function (input, index) {
						let date = moment(); // default date will shows today's date ...
						if (allInvoiceData[index] && allInvoiceData[index][`date_${index}`]) { // after change the date ...
							date = allInvoiceData[index][`date_${index}`];
						}

						let category;
						if (allInvoiceData[index] && allInvoiceData[index][`category_${index}`]) { // after change the category ...
							category = allInvoiceData[index][`category_${index}`];
						}

					//============== To Display add component ============================
						let ref = `input_${index}`;
						return (
							<div key={ref}>
								<div className="block">
									<div className="clearfix">
										<span className="pull-right" onClick={this.removeInputField.bind(this, index)} id={ref} ><button type="button" className="btn btn-danger btn-sm">Remove</button></span>
										<label className="paddingTop"> Component #{index + 1} </label>
									</div>
									<div>
										<div className="clearfix" className="setPaddingleft">

											<input
												type="text"
												className="col-lg-2 textboxSize"
												onChange={ this.handleFieldsComponent}
												id={`serial_${index}`}
												placeholder="Serial"
											/>

											<input
												type="text"
												className="col-lg-2 col-lg-offset-1 textboxSize marginRight"
												onChange={ this.handleFieldsComponent }
												id={`component_${index}`}
												placeholder="Component Name"
											/>

											<DatePicker
												className="col-lg-2 col-lg-offset-1 textboxSize"
												id={`warrantyDate_${ index }`}
												dateFormat="DD/MM/YYYY"
												selected={ date }
												peekNextMonth
												showYearDropdown
												showMonthDropdown
												dropdownMode="select"
												onChange={this.handleDateChange.bind(this, index)}
											/>

										</div>
										<br/>
										<div className="clearfix">
											<Select
												className="col-lg-3 pull-left"
												style={selectcss}
												id={`warrantyDate_${ index }`}
												value={ category }
												placeholder="Category"
												options={ categories }
												onChange={ this.handleCategoryChange.bind(this, index) }
												clearable={false}
												/>
											<div>

												<textarea
													className="textAreaSize1 col-lg-3"
													name="Address"
													id={`description_${index}`}
													onChange={ this.handleFieldsComponent }
													placeholder="Description"
												/>

											</div>
										</div>
									</div>
									<br/>
								</div>
								<br/>
							</div>
						)
					//====================================================================

					}.bind(this))}
				</div>

				</div>
			</div>
		)
	}
}

export default reduxForm({
	form: 'Edit_invoice',// a unique identifier for this form
})(Edit_invoice)
