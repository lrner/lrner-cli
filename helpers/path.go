package helpers

import (
  "fmt"
  "io/ioutil"
  "encoding/json"
  "log"
  "os"
  "strings"
  "gopkg.in/yaml.v2"


)

type Path struct {
  Path string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func (p *Path) isYAML() bool {
  return strings.HasSuffix(p.Path, "yaml") || strings.HasSuffix(p.Path, "yml")
}

func (p *Path) isJSON() bool {
  return strings.HasSuffix(p.Path, "json")
}

func (p *Path) Summary() {
  if p.SkillYAMLExists(p.Path) {
    log.Println("Importing 1 Skill from ")
  } else {
    files, err := ioutil.ReadDir("./")
    if err != nil {
      log.Fatal(err)
    }

    for _, f := range files {
      if p.SkillYAMLExists(fmt.Sprintf("%s/%s", p.Path, f.Name())) {
        fmt.Printf("Importing 1 skill from %s/%s", p.Path, f.Name())
      }
    }
  }
}

func (p *Path) IsDirectory() bool {
  fi, err := os.Stat(p.Path)
  check(err)

  switch mode := fi.Mode(); {
    case mode.IsDir():
        return true
    case mode.IsRegular():
        return false
  }
  return false

}

func (p *Path) EachJSON() []string {
  var retval []string

  if p.SkillYAMLExists(p.Path) {
    fmt.Println("Importing 1 Skill from ")
  } else {
    files, err := ioutil.ReadDir("./")
    if err != nil {
      log.Fatal(err)
    }


    for _, f := range files {
      if p.SkillYAMLExists(fmt.Sprintf("%s/%s", p.Path, f.Name())) {
        sp := Path{Path: fmt.Sprintf("%s/%s", p.Path, f.Name())}
        retval = append(retval, sp.GetJSON())
      }
    }
  }

  return retval
}

func (p *Path) GetJSON() string {
  if p.SkillYAMLExists(p.Path) {
    log.Println("Importing Data from YAML File.")
    return p.GetSkillJSON()
  }
  if p.isJSON() {
    log.Println("Importing Data from JSON File.")
    dat, err := ioutil.ReadFile(p.Path)
    check(err)
    return string(dat)
  }
  return ""
}

func (p *Path) SkillYAMLExists(path string) bool {
  if _, err := os.Stat(fmt.Sprintf("%s/%s", path, "skill.yml")); os.IsNotExist(err) {
    if _, err := os.Stat(fmt.Sprintf("%s/%s", path, "skill.yaml")); os.IsNotExist(err) {
      return false
    } else {
      log.Println("Found skill.yaml")
      log.Printf("%s/%s\n", path, "skill.yaml")
    }
  } else {
    log.Println("Found skill.yml")
    log.Printf("%s/%s\n", path, "skill.yml")
  }
  return true
}

func (p *Path) CategoryYAMLsExist(category_slug string) bool {
  if _, err := os.Stat(fmt.Sprintf("%s/%s.yml", p.Path, category_slug)); os.IsNotExist(err) {
    if _, err := os.Stat(fmt.Sprintf("%s/%s.yaml", p.Path, category_slug)); os.IsNotExist(err) {
      return false
    }
  }
  return true
}

type CategoryObject struct {
  Title string `yaml:"title"`
  Slug string `yaml:"slug"`
  Adjective string `yaml:"adjective"`
  Types string `yaml:"types"`
  Lead string `yaml:"lead"`
  Items map[string]string
}

type SkillObject struct {
  Title string `yaml:"title"`
  Description string `yaml:"description"`
  Slug string `yaml:"slug"`
  Adjective string `yaml:"adjective"`
  Categories map[string]CategoryObject
}

func (p *Path) GetSkillJSON() string {
  var filename string
  if _, err := os.Stat(fmt.Sprintf("%s/%s", p.Path, "skill.yml")); err == nil {
    filename = "skill.yml"
  } else if _, err := os.Stat(fmt.Sprintf("%s/%s", p.Path, "skill.yaml")); err == nil {
    filename = "skill.yaml"
  }

  buffer, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", p.Path, filename))
  check(err)
  var outFile []byte


  var skillObject SkillObject
  yaml.Unmarshal(buffer, &skillObject)


  for _, category := range skillObject.Categories {
    skillObject.Categories[category.Slug] = p.GetCategoryObject(category.Slug)
  }

  outFile, err = json.Marshal(skillObject)
  check(err)
  return fmt.Sprintf(string(outFile))
}

func (p *Path) GetCategoryObject(category_slug string) CategoryObject {
  var filename string
  if _, err := os.Stat(fmt.Sprintf("%s/%s.yml", p.Path, category_slug)); err == nil {
    filename = fmt.Sprintf("%s.yml", category_slug)
    if _, err := os.Stat(fmt.Sprintf("%s/%s.yaml", p.Path, category_slug)); err == nil {
      filename = fmt.Sprintf("%s.yaml", category_slug)
    }
  }

  log.Printf("For category import, reading file at: %s/%s", p.Path, filename)
  buffer, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", p.Path, filename))
  check(err)

  var categoryObject map[string]CategoryObject
  yaml.Unmarshal(buffer, &categoryObject)

  log.Printf("Returning category object for: %s\n", categoryObject[category_slug].Slug)
  log.Printf("%+v\n\n", categoryObject)
  return categoryObject[category_slug]
}
