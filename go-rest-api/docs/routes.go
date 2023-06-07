package docs

import model "be-training/go-rest-api/pkg/Laptop/model"

// swagger:route GET /
// Get List of all the laptops in collection
// responses:
//   200: GetPayloadRes200

// swagger:parameters GetPayloads
type GetPayloadReq struct {
	// in:body
 }
 
 // swagger:response GetPayloadRes200
 type GetPayloadRes200 struct {
	// in:body
	Body []model.LaptopStruct
 }





// swagger:route POST /
// Create a new laptop entry in collection 
// responses:
//   201: CreatePayloadRes201

//swagger:parameters CreatePayload
type CreatePayloadReq struct {
	// in:body
	Body model.LaptopStruct
	 }
	 //swagger:response CreatePayloadRes201
	 type CreatePayloadRes201 struct {
		// in:body
		Body model.SuccessResponse
	 }



// swagger:route PUT /
// Update a laptop entry in collection
// responses:
//   200: UpdatePayloadRes200

//swagger:parameters UpdatePayload
type UpdatePayloadReq struct {
	// in:body
	Body model.Update
	 }
	 //swagger:response UpdatePayloadRes200
	 type UpdatePayloadRes200 struct {
		// in:body
		Body model.SuccessResponse

	 }

// swagger:route DELETE /
// Delete a laptop entry in collection
// responses:
//   200: DeletePayloadRes200

//swagger:parameters DeletePayload
type DeletePayloadReq struct {
	// in:body
	Body model.IdResponse
	 }
	 //swagger:response DeletePayloadRes200
	 type DeletePayloadRes200 struct {
		// in:body
		Body model.SuccessResponse
	 }

// swagger:route GET /login
// Generate a token for authentication
// responses:
//   200: LoginPayloadRes200

//swagger:parameters LoginPayload
type LoginPayloadReq struct {
	// in:body
	 }
	 //swagger:response LoginPayloadRes200
	 type LoginPayloadRes200 struct {
		// in:body
	 }







