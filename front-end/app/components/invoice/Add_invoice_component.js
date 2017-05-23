import React, {Component} from 'react'
import { textarea, Field, FieldArray, reduxForm } from 'redux-form'
import {BootstrapTable, TableHeaderColumn} from 'react-bootstrap-table';
import Select from 'react-select';
import DatePicker from 'react-datepicker';
import moment from 'moment';
import { HelpBlock } from 'react-bootstrap'
import 'react-datepicker/dist/react-datepicker.css';
import 'react-select/dist/react-select.css';
import './styles.css';

class Add_invoice extends Component {
	constructor(props) {
		super(props);

		this.state = {
			inputs :[],
			data :[],
			serialNos : [],

			invoiceErr:'',
			invoicerErr:'',
			contactErr:'',
			serialErr:'',
			componentErr:'',
			category:''
		};
		this.handleSubmit = this.handleSubmit.bind(this);
		this.addInputField = this.addInputField.bind(this);
		this.removeInputField = this.removeInputField.bind(this);
		this.handleCategoryChange = this.handleCategoryChange.bind(this);
		this.handleDateChange = this.handleDateChange.bind(this);
		this.handleFieldsInvoice = this.handleFieldsInvoice.bind(this);
		this.handleInvoiceDateChange = this.handleInvoiceDateChange.bind(this);
		this.handleFieldsComponent = this.handleFieldsComponent.bind(this);
	}

	componentDidMount() {
		this.props.actions.fetchCategories(false, this.props.dispatch);
	}

	handleFieldsInvoice(event) {
		this.props.actions.setFieldForInvoice(event.target.id, event.target.value);
	}

	handleFieldsComponent(event) {
		this.props.actions.setFieldForComponent(event.target.id, event.target.value);
	}

	handleDateChange(componentId, date) {
		this.setState({startDate: date});
		let id = `date_${ componentId }`
		this.props.actions.setFieldForComponent(id, date);

	}

	handleInvoiceDateChange(date) {
		this.setState({startDate: date});
		let id = `Invoice_date`
		this.props.actions.setFieldForInvoice(id, date);
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
		if(data.data.length == 0){
			this.setState({componentErr: '*Add Atleast One Component'})
			status = false;
		}
		if (status) {
			this.props.actions.addInvoice(data)
		}
	}

	render() {
		let invoiceDate = moment();
		if(this.props.props.invoices.Invoice_date) {
			invoiceDate = moment(this.props.props.invoices.Invoice_date);
		}

		let borderStyle = {
			border: 'solid',
			borderWidth: '2px',
			padding: '20px 25px 20px 100px'
		};

		let selectcss ={
			width: '80%',
			overflow: 'scroll'
		}

		// let setPaddingleft = {
		// 	paddingLeft: '15px'
		// }

		let categories = [];
		for(let i = 0; i < this.props.props.category.AllCategories.length; i++){
			let ComponentInfo = { value : this.props.props.category.AllCategories[i].Id, label: this.props.props.category.AllCategories[i].Category };
			categories[i] = ComponentInfo;
		}
		let allInvoiceData = this.props.props.invoices.data;
		const { handleSubmit, pristine, reset, submitting } = this.props
		return (
			<div >

				<form onSubmit={ this.handleSubmit }>
					<button className="btn btn-info floatRight" >Save</button>
				</form>
				<button className="btn btn-success pull-right fixPositionForAdd" onClick={this.addInputField}> Add Component </button>

				<div>
				<h2 className="center"> Add Invoice </h2>
				<div style={borderStyle}>
					<div className="clearfix form-group">
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Invoice*</label>

							<input
								className="textboxSize"
								id="invoice"
								name="Invoice"
								onChange={ this.handleFieldsInvoice }
								placeholder="ex. 12MOUSE1811"
								type="text"
							/>

							<HelpBlock className="errFontStyle"> {this.state.invoiceErr} </HelpBlock>
						</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Invoicer*</label>

							<input
								className="textboxSize"
								id="invoicer"
								name="Invoicer"
								onChange={ this.handleFieldsInvoice }
								placeholder="ex. Jay systems"
								type="text"
							/>

							<HelpBlock className="errFontStyle"> {this.state.invoicerErr} </HelpBlock>
						</div>
					</div>
					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Invoice Date*</label>

								<DatePicker
									className="textboxSize"
									dateFormat="DD/MM/YYYY"
									dropdownMode="select"
									id="invoice_date"
									name="Invoice_date"
									onChange={this.handleInvoiceDateChange}
									peekNextMonth
									selected={ invoiceDate }
									showYearDropdown
									showMonthDropdown
								/>

						</div>
						<div className = "col-lg-2 col-lg-offset-2">
							<label >Contact*</label>

							<input
								className="textboxSize"
								type="text" name="Contact"
								id="contact"
								onChange={ this.handleFieldsInvoice }
								placeholder="ex. 9909970574"
							/>

							<HelpBlock className="errFontStyle"> {this.state.contactErr} </HelpBlock>
						</div>
					</div>
					<div className="clearfix form-group">
						<div className="col-lg-2 col-lg-offset-2">
							<label >Description</label>

							<textarea
								className="textAreaSize form-control"
								name="Description" id="description"
								onChange={ this.handleFieldsInvoice }
								placeholder="Description"
							/>

						</div>
						<div className="col-lg-2 col-lg-offset-2">
							<label >Address</label>

							<textarea
								className="textAreaSize form-control"
								name="Address"
								id="address"
								onChange={ this.handleFieldsInvoice }
								placeholder="Address"
							/>

						</div>
					</div>
				</div>
			</div>
			<br/>

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
		)
	}
}

export default reduxForm({
	form: 'Add_invoice',// a unique identifier for this form
})(Add_invoice)
