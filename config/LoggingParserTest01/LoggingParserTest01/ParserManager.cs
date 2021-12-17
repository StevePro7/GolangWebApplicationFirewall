using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace LoggingParserTest01
{
    public class ParserManager
    {
        public void Process(string fullString)
        {
            var str = fullString;

            // Grabs the group ID
            var group = Regex.Match(str, @"group = '(?<ID>\d+)'", RegexOptions.IgnoreCase)
                .Groups["ID"].Value;

            // Grabs everything inside teams parentheses
            var teams = Regex.Match(str, @"team in \((?<Teams>(\s*'[^']+'\s*,?)+)\)", RegexOptions.IgnoreCase)
                .Groups["Teams"].Value;

            // Trim and remove single quotes
            var teamsArray = teams.Split(new char[] { ',' }, StringSplitOptions.RemoveEmptyEntries)
                .Select(s =>
                {
                    var trimmed = s.Trim();
                    return trimmed.Substring(1, trimmed.Length - 2);
                }).ToArray();

            var i = 7;
        }
    }
}
