package httpadapter

import (
	"fmt"
	"regexp"
	"strings"
)

var kwedPattern = regexp.MustCompile(`^\d{1,2}\.\d{2}$`)

func validatePerson(dto Person) error {
	if strings.TrimSpace(dto.Lastname) == "" {
		return fmt.Errorf("lastname is required")
	}
	if strings.TrimSpace(dto.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(dto.Patronym) == "" {
		return fmt.Errorf("patronym is required")
	}
	if dto.LocalRNOKPP == 0 {
		return fmt.Errorf("local_rnokpp is required")
	}
	return nil
}

func validateLegalPerson(dto LegalPerson) error {
	if strings.TrimSpace(dto.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if strings.TrimSpace(dto.ShortName) == "" {
		return fmt.Errorf("shortname is required")
	}
	if dto.LocalEDRPOU == 0 {
		return fmt.Errorf("local_edrpou is required")
	}
	if strings.TrimSpace(dto.LocalKWED) != "" && !kwedPattern.MatchString(dto.LocalKWED) {
		return fmt.Errorf("local_kwed must match format like 22.22")
	}
	return nil
}

func validateContract(dto Contract) error {
	if dto.Number == 0 {
		return fmt.Errorf("number is required")
	}
	if strings.TrimSpace(dto.Type) == "" {
		return fmt.Errorf("type is required")
	}
	return nil
}
