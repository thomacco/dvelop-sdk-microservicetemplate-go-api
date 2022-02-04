/*
Microservice template

Mircoservice template with movie watchlist example

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Movie struct for Movie
type Movie struct {
	// Id of the movie
	Id string `json:"id"`
	// Title of the movie
	Name string `json:"name"`
	// Genre of the movie
	Genre string `json:"genre"`
	// Director of the movie
	Director string `json:"director"`
	Actors []string `json:"actors"`
	// Description of the movie
	Description string `json:"description"`
	// Type of the movie
	Type string `json:"type"`
	// Poster-Uri of the movie
	PosterUri *string `json:"posterUri,omitempty"`
	// Number of oscars
	Oscars *int32 `json:"oscars,omitempty"`
}

// NewMovie instantiates a new Movie object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovie(id string, name string, genre string, director string, actors []string, description string, type_ string) *Movie {
	this := Movie{}
	this.Id = id
	this.Name = name
	this.Genre = genre
	this.Director = director
	this.Actors = actors
	this.Description = description
	this.Type = type_
	return &this
}

// NewMovieWithDefaults instantiates a new Movie object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovieWithDefaults() *Movie {
	this := Movie{}
	return &this
}

// GetId returns the Id field value
func (o *Movie) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Movie) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Movie) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Movie) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Movie) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Movie) SetName(v string) {
	o.Name = v
}

// GetGenre returns the Genre field value
func (o *Movie) GetGenre() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Genre
}

// GetGenreOk returns a tuple with the Genre field value
// and a boolean to check if the value has been set.
func (o *Movie) GetGenreOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Genre, true
}

// SetGenre sets field value
func (o *Movie) SetGenre(v string) {
	o.Genre = v
}

// GetDirector returns the Director field value
func (o *Movie) GetDirector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Director
}

// GetDirectorOk returns a tuple with the Director field value
// and a boolean to check if the value has been set.
func (o *Movie) GetDirectorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Director, true
}

// SetDirector sets field value
func (o *Movie) SetDirector(v string) {
	o.Director = v
}

// GetActors returns the Actors field value
func (o *Movie) GetActors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Actors
}

// GetActorsOk returns a tuple with the Actors field value
// and a boolean to check if the value has been set.
func (o *Movie) GetActorsOk() ([]string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Actors, true
}

// SetActors sets field value
func (o *Movie) SetActors(v []string) {
	o.Actors = v
}

// GetDescription returns the Description field value
func (o *Movie) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Movie) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Movie) SetDescription(v string) {
	o.Description = v
}

// GetType returns the Type field value
func (o *Movie) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Movie) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Movie) SetType(v string) {
	o.Type = v
}

// GetPosterUri returns the PosterUri field value if set, zero value otherwise.
func (o *Movie) GetPosterUri() string {
	if o == nil || o.PosterUri == nil {
		var ret string
		return ret
	}
	return *o.PosterUri
}

// GetPosterUriOk returns a tuple with the PosterUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Movie) GetPosterUriOk() (*string, bool) {
	if o == nil || o.PosterUri == nil {
		return nil, false
	}
	return o.PosterUri, true
}

// HasPosterUri returns a boolean if a field has been set.
func (o *Movie) HasPosterUri() bool {
	if o != nil && o.PosterUri != nil {
		return true
	}

	return false
}

// SetPosterUri gets a reference to the given string and assigns it to the PosterUri field.
func (o *Movie) SetPosterUri(v string) {
	o.PosterUri = &v
}

// GetOscars returns the Oscars field value if set, zero value otherwise.
func (o *Movie) GetOscars() int32 {
	if o == nil || o.Oscars == nil {
		var ret int32
		return ret
	}
	return *o.Oscars
}

// GetOscarsOk returns a tuple with the Oscars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Movie) GetOscarsOk() (*int32, bool) {
	if o == nil || o.Oscars == nil {
		return nil, false
	}
	return o.Oscars, true
}

// HasOscars returns a boolean if a field has been set.
func (o *Movie) HasOscars() bool {
	if o != nil && o.Oscars != nil {
		return true
	}

	return false
}

// SetOscars gets a reference to the given int32 and assigns it to the Oscars field.
func (o *Movie) SetOscars(v int32) {
	o.Oscars = &v
}

func (o Movie) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["genre"] = o.Genre
	}
	if true {
		toSerialize["director"] = o.Director
	}
	if true {
		toSerialize["actors"] = o.Actors
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.PosterUri != nil {
		toSerialize["posterUri"] = o.PosterUri
	}
	if o.Oscars != nil {
		toSerialize["oscars"] = o.Oscars
	}
	return json.Marshal(toSerialize)
}

type NullableMovie struct {
	value *Movie
	isSet bool
}

func (v NullableMovie) Get() *Movie {
	return v.value
}

func (v *NullableMovie) Set(val *Movie) {
	v.value = val
	v.isSet = true
}

func (v NullableMovie) IsSet() bool {
	return v.isSet
}

func (v *NullableMovie) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovie(val *Movie) *NullableMovie {
	return &NullableMovie{value: val, isSet: true}
}

func (v NullableMovie) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovie) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


