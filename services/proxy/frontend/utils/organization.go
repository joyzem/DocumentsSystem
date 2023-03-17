package utils

import (
	"fmt"

	organizationDto "github.com/joyzem/documents/services/organization/dto"
	"github.com/levigross/grequests"
)

func GetOrganizationById(id int) (*organizationDto.OrganizationByIdResponse, error) {
	organizationsUrl := fmt.Sprintf("%s/organizations/%d", GetOrganizationsAddress(), id)
	organizationResp, err := grequests.Get(organizationsUrl, &grequests.RequestOptions{
		JSON: organizationDto.OrganizationByIdRequest{
			Id: id,
		}})
	var organization organizationDto.OrganizationByIdResponse
	organizationResp.JSON(&organization)
	return &organization, err
}

func GetOrganizations() (organizationDto.GetOrganizationsResponse, error) {
	organizationsUrl := fmt.Sprintf("%s/organizations", GetOrganizationsAddress())
	orgResp, err := grequests.Get(organizationsUrl, nil)
	var organizations organizationDto.GetOrganizationsResponse
	orgResp.JSON(&organizations)
	return organizations, err
}
