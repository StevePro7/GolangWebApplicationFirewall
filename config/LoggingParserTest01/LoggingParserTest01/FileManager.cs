using System;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LoggingParserTest01
{
    public class FileManager
    {
        public void Process(string path)
        {
            LogText = File.ReadAllText(path);
        }

        public string LogText { get; private set; }
    }
}
