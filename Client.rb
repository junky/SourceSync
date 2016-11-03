require 'rest-client'
require 'find'

@ignore_paths = []
@ignore_paths << ".sourcesync"
@ignore_paths << ".git/"
@ignore_paths << "temp/"

def send_file(path)
  RestClient.post( 'http://localhost:8000',
    :upload => File.new(path, 'r'),
    :path => path
  )
end

def get_last_modification_time
  source_sync_path = "./.sourcesync"
  return File.new(source_sync_path,"r").mtime if File.exists?(source_sync_path)
  return nil
end

def set_last_modification_time
  source_sync_path = "./.sourcesync"
  file = File.new(source_sync_path,"w")
  file.close
end

def file_modified?(path, last_mtime)
  return true if last_mtime.nil?
  return true if File.new(path,"r").mtime >= last_mtime
  return false  
end

def ignore_path?(path)
  @ignore_paths.each { |ignore_path_var|
    return true if path.start_with?("./" + ignore_path_var)
  }
  return false
end

Dir.chdir(File.expand_path(File.dirname(__FILE__)))

@run_files_update = true
ARGV.each do|a|
  @run_files_update = false if a.eql? "--reset_timestamp" 
end

last_mtime = get_last_modification_time
set_last_modification_time

if @run_files_update
  file_paths = []
  Find.find('.') do |path|
    file_paths << path if not File.directory?(path) and not ignore_path?(path) and file_modified?(path, last_mtime)
  end
  
  file_paths.each {|path|
      send_file(path)
      puts Time.now.strftime("%Y-%m-%d %H:%M:%S") + " : " + path
  }
end
