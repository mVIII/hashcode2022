package main

import (
	"fmt"
	"hashcode2022/utils"
	"math"
	"strconv"
)

type Person struct {
	Name    string
	Skills  map[string]int
	Working bool
}

type Skill struct {
	Name  string
	Level int
}

type Project struct {
	Name                 string
	Days                 int
	Score                int
	Deadline             int
	Skills               []*Skill
	Contributors         map[string]*Person
	PossibleContributors []*Person
	Working              bool
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func (p Project) evaluate() int {
	return int(math.Abs(float64(-50*p.Days - 25*len(p.Skills) + 5*p.Score)))
}

func main() {

	noProjects := 0
	noPeople := 0
	pplFound := false

	//[name][skill][level]
	var people []*Person
	var projects []*Project

	skillCount := 0
	var currPerson *Person
	var currProject *Project
	utils.Parse("/Downloads/small.txt", func(line int, words []string) {
		if line == 0 {
			noProjects = parseInt(words[1])
			noPeople = parseInt(words[0])
			return
		}

		if len(people) != noPeople {
			if skillCount == 0 {
				if currPerson != nil {
					people = append(people, currPerson)
					currPerson = nil
				}
				currPerson = &Person{
					Name:   words[0],
					Skills: make(map[string]int),
				}
				skillCount = parseInt(words[1])
				if len(people) == noPeople {
					pplFound = true
					skillCount = 0
				}

			} else if skillCount > 0 {
				currPerson.Skills[words[0]] = parseInt(words[1])
				skillCount--
			}
		}

		if len(projects) != noProjects && pplFound {
			if skillCount == 0 {
				if currProject != nil {
					projects = append(projects, currProject)
					currProject = nil
				}

				currProject = &Project{
					Name:         words[0],
					Days:         parseInt(words[1]),
					Score:        parseInt(words[2]),
					Deadline:     parseInt(words[3]),
					Skills:       []*Skill{},
					Working:      false,
					Contributors: make(map[string]*Person),
				}
				skillCount = parseInt(words[4])
			} else if skillCount > 0 {
				currProject.Skills = append(currProject.Skills, &Skill{
					Name:  words[0],
					Level: parseInt(words[1]),
				})
				skillCount--
			}

		}
	})

	projects = append(projects, currProject)

	sProjects := utils.Sort(projects, func(p *Project) int {
		fmt.Printf("%s,%v\n", p.Name, p.evaluate())
		return p.evaluate()
	})

	skillMap := make(map[string][]*Person)

	for _, p := range people {
		for _, skill := range utils.GetKeys(p.Skills) {
			skillMap[skill] = append(skillMap[skill], p)
		}
	}

	var brojects []Project

	running := true
	for running {
		for _, project := range sProjects {
			if project.Days == 0 {
				for skill, contrib := range project.Contributors {
					contrib.Working = !contrib.Working
					contrib.Skills[skill]++
				}
				project.Working = false
				if len(utils.Filter(brojects, func(p Project) bool {
					return p.Name == project.Name
				})) == 0 {
					brojects = append(brojects, *project)
				}

			}
		}

		if len(utils.Filter(projects, func(p *Project) bool {
			return p.Working
		})) == 0 {
			running = false
		}

		for _, project := range sProjects {
			if project.Working || project.Days == 0 {
				continue
			}
			for _, skill := range project.Skills {
				peopleWithSkills := skillMap[skill.Name]

				minMan := utils.Min(utils.Filter(peopleWithSkills, func(p *Person) bool {
					return p.Skills[skill.Name] >= utils.Find(project.Skills, func(s *Skill) bool {
						return s.Name == skill.Name
					}).Level && !p.Working
				}), func(p *Person) int {
					return p.Skills[skill.Name]
				})

				if minMan != nil {
					project.PossibleContributors = append(project.PossibleContributors, minMan)
				}

				if len(project.PossibleContributors) == len(project.Skills) {
					project.Working = true
					running = true
					for _, contributor := range project.PossibleContributors {
						contributor.Working = true
						project.Contributors[skill.Name] = contributor
					}
					project.PossibleContributors = []*Person{}
				}

			}

		}

		for _, project := range sProjects {
			if project.Working {
				project.Days--
			}
		}

	}

	fmt.Println(brojects)

}
