require 'yaml'
require 'json'

y = YAML.load File.read 'schema.yml'
j = JSON.pretty_generate y
File.write 'schema.json', j
