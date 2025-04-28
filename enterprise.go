package goattck

import (
	"fmt"

	"github.com/msadministrator/goattck/internal/logger"
	"github.com/msadministrator/goattck/internal/models"
)

var (
	slogger = logger.NewLogger(logger.Info)
)

// The raw representation of a custom data model used by both pyattck & goattck
type RawEnterpriseAttck struct {
	// The type of framework
	AttckType string `json:"type"`
	// THe unique ID or hash of the JSON object
	ID string `json:"id"`
	// An array of data models structs for each entity
	Objects []interface{} `json:"objects"`
	// The defined version of the framework
	SpecVersion string `json:"spec_version"`
	// The last time the data was updated/modified.
	LastUpdated string `json:"last_updated"`
	// Whether or not this is a revoked version of the framework
	Revoked bool `json:"revoked"`
}

type enterprise interface {
	New(url string) (Enterprise, error)
	Load(force bool) (Enterprise, error)
}

// Enterprise struct represents the MITRE ATT&CK Enterprise framework
type Enterprise struct {
	Actors                []*models.ActorObject
	Campaigns             []*models.CampaignObject
	Controls              []*models.ControlObject
	DataComponents        []*models.DataComponentObject
	DataSources           []*models.DataSourceObject
	Defintions            []*models.MarkingDefinitionObject
	Malwares              []*models.MalwareObject
	Matrices              []*models.MatrixObject
	Mitigations           []*models.MitigationObject
	Relationships         []*models.RelationshipObject
	Tactics               []*models.TacticObject
	Techniques            []*models.TechniqueObject
	Tools                 []*models.ToolObject
	rawData               RawEnterpriseAttck
	attackRelationshipMap map[string][]string
	url                   string
}

var _ (enterprise) = new(Enterprise)

// Checks if the data file exists locally and if not will download
// and save the latest data. If it does exist, we use it.
func (e Enterprise) New(url string) (Enterprise, error) {
	e.url = url
	s := Storage{}

	ok, err := s.ExistsLocally()
	if err != nil {
		slogger.Warning(fmt.Sprintf("error checking for file locally: %s", err))
	}
	if !ok {
		// If it doesn't exist locally, we download it
		slogger.Debug(fmt.Sprintf("downloading and saving latest json data from %s", url))
		e.rawData, err = s.DownloadAndSave(url)
		if err != nil {
			slogger.Fatal(fmt.Sprintf("unable to download from url: %s %s", url, err))
		}
	} else {
		slogger.Debug("retrieving json from local disk")
		e.rawData, err = s.Retrieve()
		if err != nil {
			slogger.Fatal(fmt.Sprintf("error retrieving file from disk: %s", err))
		}
	}

	return e, nil
}

// Loads our data models from our downloaded data
func (e Enterprise) Load(force bool) (Enterprise, error) {
	err := e.loadDataModels()
	if err != nil {
		slogger.Error(fmt.Sprintf("Error loading data models: %s", err))
		return e, err
	}
	return e, err
}

// Load data models from rawEnterpriseAttck struct
func (e *Enterprise) loadDataModels() error {
	e.attackRelationshipMap = map[string][]string{}
	for _, value := range e.rawData.Objects {
		v, ok := value.(map[string]interface{})
		if !ok {
			slogger.Error("error casting value to map")
		}
		switch v["type"] {
		case "intrusion-set":
			actor, err := models.NewActor(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating actor: %s", err))
			}
			e.Actors = append(e.Actors, actor)
		case "campaign":
			campaign, err := models.NewCampaign(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating campaign: %s", err))
			}
			e.Campaigns = append(e.Campaigns, campaign)
		case "x-mitre-data-component":
			dataComponent, err := models.NewDataComponent(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating data component: %s", err))
			}
			e.DataComponents = append(e.DataComponents, dataComponent)
		case "x-mitre-data-source":
			dataSource, err := models.NewDataSource(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating data source: %s", err))
			}
			e.DataSources = append(e.DataSources, dataSource)
		case "marking-definition":
			markingDefinition, err := models.NewMarkingDefinition(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating marking definition: %s", err))
			}
			e.Defintions = append(e.Defintions, markingDefinition)
		case "malware":
			malware, err := models.NewMalware(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating malware: %s", err))
			}
			e.Malwares = append(e.Malwares, malware)
		case "course-of-action":
			mitigation, err := models.NewMitigation(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating mitigation: %s", err))
			}
			e.Mitigations = append(e.Mitigations, mitigation)
		case "x-mitre-matrix":
			matrix, err := models.NewMatrix(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating matrix: %s", err))
			}
			e.Matrices = append(e.Matrices, matrix)
		case "relationship":
			relationship, err := models.NewRelationship(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating relationship: %s", err))
			}
			e.Relationships = append(e.Relationships, relationship)
			if relationship.RelationshipType != "revoked-by" {
				sourceID := relationship.SourceRef
				targetID := relationship.TargetRef
				if _, ok := e.attackRelationshipMap[sourceID]; !ok {
					e.attackRelationshipMap[sourceID] = []string{}
				}
				if _, ok := e.attackRelationshipMap[targetID]; !ok {
					e.attackRelationshipMap[targetID] = []string{}
				}
				found := false
				for _, val := range e.attackRelationshipMap[sourceID] {
					if val == targetID {
						found = true
					}
				}
				if !found {
					e.attackRelationshipMap[sourceID] = append(e.attackRelationshipMap[sourceID], targetID)
				}
				found = false
				for _, val := range e.attackRelationshipMap[targetID] {
					if val == sourceID {
						found = true
					}
				}
				if !found {
					e.attackRelationshipMap[targetID] = append(e.attackRelationshipMap[targetID], sourceID)
				}
			}
		case "x-mitre-tactic":
			tactic, err := models.NewTactic(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating tactic: %s", err))
			}
			e.Tactics = append(e.Tactics, tactic)
		case "attack-pattern":
			technique, err := models.NewTechnique(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating technique: %s", err))
			}
			e.Techniques = append(e.Techniques, technique)
		case "tool":
			tool, err := models.NewTool(v)
			if err != nil {
				slogger.Error(fmt.Sprintf("Error creating tool: %s", err))
			}
			e.Tools = append(e.Tools, tool)
		}
	}
	e.setRelationships()
	return nil
}

// This method is used once we have collected our structs
// to build relationships by collecting structs as properties
// on the different entities
func (e Enterprise) setRelationships() {
	e.setActorRelationships()
	e.setCampaignRelationships()
	e.setDataComponentRelationships()
	e.setDataSourceRelationships()
	e.setMalwareRelationships()
	e.setMitigationRelationships()
	e.setTacticRelationships()
	e.setTechniqueRelationships()
	e.setToolRelationships()
}

func (e Enterprise) setActorRelationships() {
	for _, actor := range e.Actors {
		if e.attackRelationshipMap[actor.Id] != nil {
			var malwares []*models.MalwareObject
			var tools []*models.ToolObject
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[actor.Id] {
				for _, malware := range e.Malwares {
					if malware.Id == relationshipId {
						malwares = append(malwares, malware)
					}
				}
				for _, tool := range e.Tools {
					if tool.Id == relationshipId {
						tools = append(tools, tool)
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			actor.Malwares = malwares
			actor.Tools = tools
			actor.Techniques = techniques
		}
	}
}

func (e Enterprise) setCampaignRelationships() {
	for _, campaign := range e.Campaigns {
		if e.attackRelationshipMap[campaign.Id] != nil {
			var malwares []*models.MalwareObject
			var tools []*models.ToolObject
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[campaign.Id] {
				for _, malware := range e.Malwares {
					if malware.Id == relationshipId {
						malwares = append(malwares, malware)
					}
				}
				for _, tool := range e.Tools {
					if tool.Id == relationshipId {
						tools = append(tools, tool)
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			campaign.Malwares = malwares
			campaign.Tools = tools
			campaign.Techniques = techniques
		}
	}
}

func (e Enterprise) setDataComponentRelationships() {
	for _, dataComponent := range e.DataComponents {
		if e.attackRelationshipMap[dataComponent.Id] != nil {
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[dataComponent.Id] {
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			dataComponent.Techniques = techniques
		}
	}
}

func (e Enterprise) setDataSourceRelationships() {
	for _, dataSource := range e.DataSources {
		if e.attackRelationshipMap[dataSource.Id] != nil {
			var dataComponents []*models.DataComponentObject
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[dataSource.Id] {
				for _, dataComponent := range e.DataComponents {
					if dataComponent.Id == relationshipId {
						dataComponents = append(dataComponents, dataComponent)
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			dataSource.DataComponents = dataComponents
			dataSource.Techniques = techniques
		}
	}
}

func (e Enterprise) setMalwareRelationships() {
	for _, malware := range e.Malwares {
		if e.attackRelationshipMap[malware.Id] != nil {
			var actors []*models.ActorObject
			var campaigns []*models.CampaignObject
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[malware.Id] {
				for _, actor := range e.Actors {
					if actor.Id == relationshipId {
						actors = append(actors, actor)
					}
				}
				for _, campaign := range e.Campaigns {
					if campaign.Id == relationshipId {
						campaigns = append(campaigns, campaign)
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			malware.Actors = actors
			malware.Campaigns = campaigns
			malware.Techniques = techniques
		}
	}
}

func (e Enterprise) setMitigationRelationships() {
	for _, mitigation := range e.Mitigations {
		var techniques []*models.TechniqueObject
		if e.attackRelationshipMap[mitigation.Id] != nil {
			for _, relationshipId := range e.attackRelationshipMap[mitigation.Id] {
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
		}
		if len(mitigation.ExternalReferences) > 0 {
			for _, rel := range mitigation.ExternalReferences {
				if rel.SourceName == "mitre-attack" {
					for _, technique := range e.Techniques {
						if technique.GetExternalID() == rel.ExternalId {
							techniques = append(techniques, technique)
						}
					}
				}
			}
		}
		mitigation.Techniques = techniques
	}
}

func (e Enterprise) setTacticRelationships() {
	for _, tactic := range e.Tactics {
		var techniques []*models.TechniqueObject
		for _, technique := range e.Techniques {
			if technique.KillChainPhases != nil {
				for _, phase := range technique.KillChainPhases {
					if phase.PhaseName == tactic.XMitreShortname {
						techniques = append(techniques, technique)
					}
				}
			}
		}
		tactic.Techniques = techniques
	}
}

func (e Enterprise) setTechniqueRelationships() {
	for _, technique := range e.Techniques {
		if e.attackRelationshipMap[technique.Id] != nil {
			var actors []*models.ActorObject
			var campaigns []*models.CampaignObject
			var dataComponents []*models.DataComponentObject
			var dataSources []*models.DataSourceObject
			var malwares []*models.MalwareObject
			var mitigations []*models.MitigationObject
			var tactics []*models.TacticObject
			var techniques []*models.TechniqueObject
			var tools []*models.ToolObject
			for _, relationshipId := range e.attackRelationshipMap[technique.Id] {
				for _, actor := range e.Actors {
					if actor.Id == relationshipId {
						actors = append(actors, actor)
					}
				}
				for _, campaign := range e.Campaigns {
					if campaign.Id == relationshipId {
						campaigns = append(campaigns, campaign)
					}
				}
				for _, dataComponent := range e.DataComponents {
					if dataComponent.Id == relationshipId {
						dataComponents = append(dataComponents, dataComponent)
					}
				}
				for _, dataSource := range e.DataSources {
					if dataSource.Id == relationshipId {
						dataSources = append(dataSources, dataSource)
					}
				}
				for _, malware := range e.Malwares {
					if malware.Id == relationshipId {
						malwares = append(malwares, malware)
					}
				}
				for _, mitigation := range e.Mitigations {
					if mitigation.Id == relationshipId {
						mitigations = append(mitigations, mitigation)
					}
				}
				for _, phase := range technique.KillChainPhases {
					for _, tactic := range e.Tactics {
						if tactic.XMitreShortname == phase.KillChainName {
							tactics = append(tactics, tactic)
						}
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
				for _, tool := range e.Tools {
					if tool.Id == relationshipId {
						tools = append(tools, tool)
					}
				}
			}
			technique.Actors = actors
			technique.Campaigns = campaigns
			technique.DataComponents = dataComponents
			technique.DataSources = dataSources
			technique.Malwares = malwares
			technique.Mitigations = mitigations
			technique.Tactics = tactics
			technique.Techniques = techniques
			technique.Tools = tools
		}
	}
}

func (e Enterprise) setToolRelationships() {
	for _, tool := range e.Tools {
		if e.attackRelationshipMap[tool.Id] != nil {
			var actors []*models.ActorObject
			var campaigns []*models.CampaignObject
			var techniques []*models.TechniqueObject
			for _, relationshipId := range e.attackRelationshipMap[tool.Id] {
				for _, actor := range e.Actors {
					if actor.Id == relationshipId {
						actors = append(actors, actor)
					}
				}
				for _, campaign := range e.Campaigns {
					if campaign.Id == relationshipId {
						campaigns = append(campaigns, campaign)
					}
				}
				for _, technique := range e.Techniques {
					if technique.Id == relationshipId {
						techniques = append(techniques, technique)
					}
				}
			}
			tool.Actors = actors
			tool.Campaigns = campaigns
			tool.Techniques = techniques
		}
	}
}
